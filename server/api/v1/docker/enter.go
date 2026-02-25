package docker

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DockerClusterApi
	ComputeNodeApi
}

var (
	dockerClusterService = service.ServiceGroupApp.DockerServiceGroup.DockerClusterService
	computeNodeService   = service.ServiceGroupApp.DockerServiceGroup.ComputeNodeService
)
