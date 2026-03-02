package dockerCluster

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type initDockerCluster struct{}

const initOrderDockerCluster = system.InitOrderSystem + 10

// auto run
func init() {
	system.RegisterInit(initOrderDockerCluster, &initDockerCluster{})
}

func (i *initDockerCluster) InitializerName() string {
	return "docker_cluster_init"
}

func (i *initDockerCluster) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initDockerCluster) TableCreated(ctx context.Context) bool {
	return true
}

func (i *initDockerCluster) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	// 检查是否已初始化
	if i.DataInserted(ctx) {
		return ctx, nil
	}

	// 1. 创建API
	apis := []sysModel.SysApi{
		{ApiGroup: "Docker集群管理", Method: "POST", Path: "/dockerCluster/createDockerCluster", Description: "创建Docker集群"},
		{ApiGroup: "Docker集群管理", Method: "DELETE", Path: "/dockerCluster/deleteDockerCluster", Description: "删除Docker集群"},
		{ApiGroup: "Docker集群管理", Method: "DELETE", Path: "/dockerCluster/deleteDockerClusterByIds", Description: "批量删除Docker集群"},
		{ApiGroup: "Docker集群管理", Method: "PUT", Path: "/dockerCluster/updateDockerCluster", Description: "更新Docker集群"},
		{ApiGroup: "Docker集群管理", Method: "GET", Path: "/dockerCluster/findDockerCluster", Description: "根据ID获取Docker集群"},
		{ApiGroup: "Docker集群管理", Method: "GET", Path: "/dockerCluster/getDockerClusterList", Description: "获取Docker集群列表"},
		{ApiGroup: "Docker集群管理", Method: "GET", Path: "/dockerCluster/getDockerClusterCredentials", Description: "获取Docker集群凭证"},
		{ApiGroup: "Docker集群管理", Method: "GET", Path: "/dockerCluster/getAllDockerClusters", Description: "获取所有Docker集群"},
	}

	for _, api := range apis {
		var existing sysModel.SysApi
		if err := db.Where("path = ? AND method = ?", api.Path, api.Method).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&api).Error; err != nil {
				global.GVA_LOG.Error("创建API失败", zap.Error(err))
			}
		}
	}

	// 2. 创建菜单
	menu := sysModel.SysBaseMenu{
		MenuLevel: 0,
		ParentId:  0,
		Path:      "dockerCluster",
		Name:      "dockerCluster",
		Hidden:    false,
		Component: "view/dockerCluster/dockerCluster.vue",
		Sort:      100,
		Meta: sysModel.Meta{
			Title:     "Docker集群管理",
			Icon:      "cpu",
			KeepAlive: true,
		},
	}

	// 检查菜单是否已存在
	var existingMenu sysModel.SysBaseMenu
	if err := db.Where("path = ?", menu.Path).First(&existingMenu).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(&menu).Error; err != nil {
			return ctx, errors.Wrap(err, "创建Docker集群菜单失败!")
		}

		// 获取刚创建的菜单ID
		var newMenu sysModel.SysBaseMenu
		db.Where("path = ?", menu.Path).First(&newMenu)

		// 3. 创建按钮权限
		buttons := []sysModel.SysBaseMenuBtn{
			{Name: "新增", Desc: "dockerCluster_btn_add", SysBaseMenuID: newMenu.ID},
			{Name: "删除", Desc: "dockerCluster_btn_delete", SysBaseMenuID: newMenu.ID},
			{Name: "编辑", Desc: "dockerCluster_btn_edit", SysBaseMenuID: newMenu.ID},
			{Name: "查看凭证", Desc: "dockerCluster_btn_credentials", SysBaseMenuID: newMenu.ID},
		}

		for _, btn := range buttons {
			if err := db.Create(&btn).Error; err != nil {
				global.GVA_LOG.Error("创建按钮权限失败", zap.Error(err))
			}
		}

		// 4. 授权给超级管理员角色(888)
		// 菜单授权
		db.Exec("INSERT INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) VALUES (?, ?)", 888, newMenu.ID)

		// 按钮授权
		for _, btn := range buttons {
			db.Exec("INSERT INTO sys_authority_btns (authority_id, menu_id, sys_base_menu_btn_id) SELECT 888, ?, id FROM sys_base_menu_btns WHERE sys_base_menu_id = ? AND `desc` = ?", newMenu.ID, newMenu.ID, btn.Desc)
		}

		// 5. 创建Casbin权限规则
		casbinRules := []struct {
			Path   string
			Method string
		}{
			{"/dockerCluster/createDockerCluster", "POST"},
			{"/dockerCluster/deleteDockerCluster", "DELETE"},
			{"/dockerCluster/deleteDockerClusterByIds", "DELETE"},
			{"/dockerCluster/updateDockerCluster", "PUT"},
			{"/dockerCluster/findDockerCluster", "GET"},
			{"/dockerCluster/getDockerClusterList", "GET"},
			{"/dockerCluster/getDockerClusterCredentials", "GET"},
			{"/dockerCluster/getAllDockerClusters", "GET"},
		}

		for _, rule := range casbinRules {
			db.Exec("INSERT INTO casbin_rule (p_type, v0, v1, v2) VALUES (?, ?, ?, ?)", "p", "888", rule.Path, rule.Method)
		}
	}

	next := context.WithValue(ctx, i.InitializerName(), true)
	return next, nil
}

func (i *initDockerCluster) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}

	// 检查菜单是否已存在
	var menu sysModel.SysBaseMenu
	if err := db.Where("path = ?", "dockerCluster").First(&menu).Error; err != nil {
		return false
	}

	return true
}
