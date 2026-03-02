package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	{
		initRouter.POST("initdb", dbApi.InitDB)                 // 初始化数据库
		initRouter.POST("checkdb", dbApi.CheckDB)               // 检测是否需要初始化数据库
		initRouter.POST("initDockerCluster", initDockerCluster) // 初始化Docker集群菜单和API
	}
}

// initDockerCluster 初始化Docker集群菜单和API
func initDockerCluster(c *gin.Context) {
	// 使用全局数据库连接
	db := global.GVA_DB

	// 1. 创建API
	apis := []struct {
		ApiGroup    string
		Method      string
		Path        string
		Description string
	}{
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
		db.Exec("INSERT IGNORE INTO sys_apis (created_at, updated_at, path, description, api_group, method) VALUES (NOW(), NOW(), ?, ?, ?, ?)", api.Path, api.Description, api.ApiGroup, api.Method)
	}

	// 2. 创建菜单
	var menuId uint
	result := db.Raw("INSERT INTO sys_base_menus (created_at, updated_at, menu_level, parent_id, path, name, hidden, component, sort, keep_alive, default_menu, title, icon, close_tab) SELECT NOW(), NOW(), 0, 0, 'dockerCluster', 'dockerCluster', 0, 'view/dockerCluster/dockerCluster.vue', 100, 1, 0, 'Docker集群管理', 'cpu', 0 FROM DUAL WHERE NOT EXISTS (SELECT 1 FROM sys_base_menus WHERE path = 'dockerCluster')").Scan(&menuId)
	if result.Error != nil {
		global.GVA_LOG.Error("创建菜单失败", zap.Error(result.Error))
	}

	// 获取菜单ID
	db.Raw("SELECT id FROM sys_base_menus WHERE path = ?", "dockerCluster").Scan(&menuId)

	if menuId > 0 {
		// 3. 创建按钮权限
		buttons := []struct {
			Name string
			Desc string
		}{
			{Name: "新增", Desc: "dockerCluster_btn_add"},
			{Name: "删除", Desc: "dockerCluster_btn_delete"},
			{Name: "编辑", Desc: "dockerCluster_btn_edit"},
			{Name: "查看凭证", Desc: "dockerCluster_btn_credentials"},
		}

		for _, btn := range buttons {
			db.Exec("INSERT IGNORE INTO sys_base_menu_btns (created_at, updated_at, name, desc, sys_base_menu_id) VALUES (NOW(), NOW(), ?, ?, ?)", btn.Name, btn.Desc, menuId)
		}

		// 4. 授权给超级管理员角色(888) - 菜单
		db.Exec("INSERT IGNORE INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) VALUES (?, ?)", 888, menuId)

		// 5. 授权按钮
		for _, btn := range buttons {
			var btnId uint
			db.Raw("SELECT id FROM sys_base_menu_btns WHERE sys_base_menu_id = ? AND `desc` = ?", menuId, btn.Desc).Scan(&btnId)
			if btnId > 0 {
				db.Exec("INSERT IGNORE INTO sys_authority_btns (authority_id, menu_id, sys_base_menu_btn_id) VALUES (?, ?, ?)", 888, menuId, btnId)
			}
		}

		// 6. 创建Casbin权限规则
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
			db.Exec("INSERT IGNORE INTO casbin_rule (p_type, v0, v1, v2) VALUES (?, ?, ?, ?)", "p", "888", rule.Path, rule.Method)
		}
	}

	c.JSON(200, gin.H{"code": 0, "msg": "Docker集群菜单和API初始化成功"})
}
