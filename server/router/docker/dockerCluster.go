package docker

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DockerClusterRouter struct {}

// InitDockerClusterRouter 初始化 Docker集群管理 路由信息
func (s *DockerClusterRouter) InitDockerClusterRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dockerClusterRouter := Router.Group("dockerCluster").Use(middleware.OperationRecord())
	dockerClusterRouterWithoutRecord := Router.Group("dockerCluster")
	dockerClusterRouterWithoutAuth := PublicRouter.Group("dockerCluster")
	{
		dockerClusterRouter.POST("createDockerCluster", dockerClusterApi.CreateDockerCluster)   // 新建Docker集群管理
		dockerClusterRouter.DELETE("deleteDockerCluster", dockerClusterApi.DeleteDockerCluster) // 删除Docker集群管理
		dockerClusterRouter.DELETE("deleteDockerClusterByIds", dockerClusterApi.DeleteDockerClusterByIds) // 批量删除Docker集群管理
		dockerClusterRouter.PUT("updateDockerCluster", dockerClusterApi.UpdateDockerCluster)    // 更新Docker集群管理
	}
	{
		dockerClusterRouterWithoutRecord.GET("findDockerCluster", dockerClusterApi.FindDockerCluster)        // 根据ID获取Docker集群管理
		dockerClusterRouterWithoutRecord.GET("getDockerClusterList", dockerClusterApi.GetDockerClusterList)  // 获取Docker集群管理列表
	}
	{
	    dockerClusterRouterWithoutAuth.GET("getDockerClusterPublic", dockerClusterApi.GetDockerClusterPublic)  // Docker集群管理开放接口
	}
}
