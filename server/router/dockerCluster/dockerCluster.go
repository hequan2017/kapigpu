package dockerCluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DockerClusterRouter struct {
}

// InitDockerClusterRouter 初始化 DockerCluster 路由信息
func (s *DockerClusterRouter) InitDockerClusterRouter(Router *gin.RouterGroup) {
	dockerClusterRouter := Router.Group("dockerCluster").Use(middleware.OperationRecord())
	dockerClusterRouterWithoutRecord := Router.Group("dockerCluster")
	{
		dockerClusterRouter.POST("createDockerCluster", v1.ApiGroupApp.DockerClusterApi.CreateDockerCluster)             // 新建DockerCluster
		dockerClusterRouter.DELETE("deleteDockerCluster", v1.ApiGroupApp.DockerClusterApi.DeleteDockerCluster)         // 删除DockerCluster
		dockerClusterRouter.DELETE("deleteDockerClusterByIds", v1.ApiGroupApp.DockerClusterApi.DeleteDockerClusterByIds) // 批量删除DockerCluster
		dockerClusterRouter.PUT("updateDockerCluster", v1.ApiGroupApp.DockerClusterApi.UpdateDockerCluster)              // 更新DockerCluster
	}
	{
		dockerClusterRouterWithoutRecord.GET("findDockerCluster", v1.ApiGroupApp.DockerClusterApi.FindDockerCluster)                // 根据ID获取DockerCluster
		dockerClusterRouterWithoutRecord.GET("getDockerClusterList", v1.ApiGroupApp.DockerClusterApi.GetDockerClusterList)          // 获取DockerCluster列表
		dockerClusterRouterWithoutRecord.GET("getDockerClusterCredentials", v1.ApiGroupApp.DockerClusterApi.GetDockerClusterCredentials) // 获取集群凭证
		dockerClusterRouterWithoutRecord.GET("getAllDockerClusters", v1.ApiGroupApp.DockerClusterApi.GetAllDockerClusters)          // 获取所有集群（下拉选择）
	}
}
