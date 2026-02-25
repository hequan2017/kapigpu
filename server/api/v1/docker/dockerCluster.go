package docker

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/docker"
    dockerReq "github.com/flipped-aurora/gin-vue-admin/server/model/docker/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DockerClusterApi struct {}



// CreateDockerCluster 创建Docker集群管理
// @Tags DockerCluster
// @Summary 创建Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body docker.DockerCluster true "创建Docker集群管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dockerCluster/createDockerCluster [post]
func (dockerClusterApi *DockerClusterApi) CreateDockerCluster(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var dockerCluster docker.DockerCluster
	err := c.ShouldBindJSON(&dockerCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dockerClusterService.CreateDockerCluster(ctx,&dockerCluster)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDockerCluster 删除Docker集群管理
// @Tags DockerCluster
// @Summary 删除Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body docker.DockerCluster true "删除Docker集群管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dockerCluster/deleteDockerCluster [delete]
func (dockerClusterApi *DockerClusterApi) DeleteDockerCluster(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := dockerClusterService.DeleteDockerCluster(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDockerClusterByIds 批量删除Docker集群管理
// @Tags DockerCluster
// @Summary 批量删除Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dockerCluster/deleteDockerClusterByIds [delete]
func (dockerClusterApi *DockerClusterApi) DeleteDockerClusterByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := dockerClusterService.DeleteDockerClusterByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDockerCluster 更新Docker集群管理
// @Tags DockerCluster
// @Summary 更新Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body docker.DockerCluster true "更新Docker集群管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dockerCluster/updateDockerCluster [put]
func (dockerClusterApi *DockerClusterApi) UpdateDockerCluster(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var dockerCluster docker.DockerCluster
	err := c.ShouldBindJSON(&dockerCluster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dockerClusterService.UpdateDockerCluster(ctx,dockerCluster)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDockerCluster 用id查询Docker集群管理
// @Tags DockerCluster
// @Summary 用id查询Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Docker集群管理"
// @Success 200 {object} response.Response{data=docker.DockerCluster,msg=string} "查询成功"
// @Router /dockerCluster/findDockerCluster [get]
func (dockerClusterApi *DockerClusterApi) FindDockerCluster(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redockerCluster, err := dockerClusterService.GetDockerCluster(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redockerCluster, c)
}
// GetDockerClusterList 分页获取Docker集群管理列表
// @Tags DockerCluster
// @Summary 分页获取Docker集群管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query dockerReq.DockerClusterSearch true "分页获取Docker集群管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dockerCluster/getDockerClusterList [get]
func (dockerClusterApi *DockerClusterApi) GetDockerClusterList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo dockerReq.DockerClusterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dockerClusterService.GetDockerClusterInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetDockerClusterPublic 不需要鉴权的Docker集群管理接口
// @Tags DockerCluster
// @Summary 不需要鉴权的Docker集群管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dockerCluster/getDockerClusterPublic [get]
func (dockerClusterApi *DockerClusterApi) GetDockerClusterPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    dockerClusterService.GetDockerClusterPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的Docker集群管理接口信息",
    }, "获取成功", c)
}
