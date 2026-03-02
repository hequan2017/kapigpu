package dockerCluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dockerCluster"
	dockerClusterReq "github.com/flipped-aurora/gin-vue-admin/server/model/dockerCluster/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DockerClusterApi struct{}

var dockerClusterService = service.ServiceGroupApp.DockerClusterService

// CreateDockerCluster 创建Docker集群
// @Tags     DockerCluster
// @Summary  创建Docker集群
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     dockerCluster.DockerCluster true "Docker集群信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /dockerCluster/createDockerCluster [post]
func (dockerClusterApi *DockerClusterApi) CreateDockerCluster(c *gin.Context) {
	var dc dockerCluster.DockerCluster
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证必填字段
	if dc.ClusterName == "" || dc.CaCert == "" || dc.ClientCert == "" || dc.ClientKey == "" {
		response.FailWithMessage("集群名称、CA证书、客户端证书、客户端私钥不能为空", c)
		return
	}

	if err := dockerClusterService.CreateDockerCluster(&dc); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败："+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDockerCluster 删除Docker集群
// @Tags     DockerCluster
// @Summary  删除Docker集群
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     dockerCluster.DockerCluster true "ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /dockerCluster/deleteDockerCluster [delete]
func (dockerClusterApi *DockerClusterApi) DeleteDockerCluster(c *gin.Context) {
	var dc dockerCluster.DockerCluster
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dockerClusterService.DeleteDockerCluster(dc); err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败："+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDockerClusterByIds 批量删除Docker集群
// @Tags     DockerCluster
// @Summary  批量删除Docker集群
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.IdsReq true "ID列表"
// @Success  200  {object} response.Response{msg=string} "批量删除成功"
// @Router   /dockerCluster/deleteDockerClusterByIds [delete]
func (dockerClusterApi *DockerClusterApi) DeleteDockerClusterByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dockerClusterService.DeleteDockerClusterByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败", zap.Error(err))
		response.FailWithMessage("批量删除失败："+err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDockerCluster 更新Docker集群
// @Tags     DockerCluster
// @Summary  更新Docker集群
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     dockerCluster.DockerCluster true "Docker集群信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /dockerCluster/updateDockerCluster [put]
func (dockerClusterApi *DockerClusterApi) UpdateDockerCluster(c *gin.Context) {
	var dc dockerCluster.DockerCluster
	err := c.ShouldBindJSON(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if dc.ID == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	if err := dockerClusterService.UpdateDockerCluster(dc); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败："+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDockerCluster 用ID查询Docker集群
// @Tags     DockerCluster
// @Summary  用ID查询Docker集群
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data query     dockerCluster.DockerCluster true "ID"
// @Success  200  {object} response.Response{data=dockerCluster.DockerCluster,msg=string} "查询成功"
// @Router   /dockerCluster/findDockerCluster [get]
func (dockerClusterApi *DockerClusterApi) FindDockerCluster(c *gin.Context) {
	var dc dockerCluster.DockerCluster
	err := c.ShouldBindQuery(&dc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if dc.ID == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	if reDc, err := dockerClusterService.GetDockerCluster(dc.ID); err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败："+err.Error(), c)
	} else {
		response.OkWithDetailed(reDc, "查询成功", c)
	}
}

// GetDockerClusterList 分页获取Docker集群列表
// @Tags     DockerCluster
// @Summary  分页获取Docker集群列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data query     dockerClusterReq.DockerClusterSearch true "分页参数"
// @Success  200  {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router   /dockerCluster/getDockerClusterList [get]
func (dockerClusterApi *DockerClusterApi) GetDockerClusterList(c *gin.Context) {
	var pageInfo dockerClusterReq.DockerClusterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := dockerClusterService.GetDockerClusterInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败："+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetDockerClusterCredentials 获取集群凭证（用于连接Docker集群）
// @Tags     DockerCluster
// @Summary  获取集群凭证（用于连接Docker集群）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id query     uint true "集群ID"
// @Success  200  {object} response.Response{data=dockerCluster.DockerCluster,msg=string} "获取成功"
// @Router   /dockerCluster/getDockerClusterCredentials [get]
func (dockerClusterApi *DockerClusterApi) GetDockerClusterCredentials(c *gin.Context) {
	var req struct {
		ID uint `json:"id" form:"id"`
	}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.ID == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}

	if dc, err := dockerClusterService.GetDockerClusterCredentials(req.ID); err != nil {
		global.GVA_LOG.Error("获取凭证失败", zap.Error(err))
		response.FailWithMessage("获取凭证失败："+err.Error(), c)
	} else {
		response.OkWithDetailed(dc, "获取成功", c)
	}
}

// GetAllDockerClusters 获取所有集群（用于下拉选择）
// @Tags     DockerCluster
// @Summary  获取所有集群（用于下拉选择）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]dockerCluster.DockerCluster,msg=string} "获取成功"
// @Router   /dockerCluster/getAllDockerClusters [get]
func (dockerClusterApi *DockerClusterApi) GetAllDockerClusters(c *gin.Context) {
	if list, err := dockerClusterService.GetAllDockerClusters(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败："+err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
