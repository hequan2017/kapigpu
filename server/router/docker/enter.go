package docker

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	DockerClusterRouter
	ComputeNodeRouter
}

var (
	dockerClusterApi = api.ApiGroupApp.DockerApiGroup.DockerClusterApi
	computeNodeApi   = api.ApiGroupApp.DockerApiGroup.ComputeNodeApi
)
