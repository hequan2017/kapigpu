import service from '@/utils/request'

// @Tags DockerCluster
// @Summary 创建DockerCluster
export const createDockerCluster = (data) => {
  return service({
    url: '/dockerCluster/createDockerCluster',
    method: 'post',
    data
  })
}

// @Tags DockerCluster
// @Summary 删除DockerCluster
export const deleteDockerCluster = (data) => {
  return service({
    url: '/dockerCluster/deleteDockerCluster',
    method: 'delete',
    data
  })
}

// @Tags DockerCluster
// @Summary 批量删除DockerCluster
export const deleteDockerClusterByIds = (data) => {
  return service({
    url: '/dockerCluster/deleteDockerClusterByIds',
    method: 'delete',
    data
  })
}

// @Tags DockerCluster
// @Summary 更新DockerCluster
export const updateDockerCluster = (data) => {
  return service({
    url: '/dockerCluster/updateDockerCluster',
    method: 'put',
    data
  })
}

// @Tags DockerCluster
// @Summary 用id查询DockerCluster
export const findDockerCluster = (params) => {
  return service({
    url: '/dockerCluster/findDockerCluster',
    method: 'get',
    params
  })
}

// @Tags DockerCluster
// @Summary 分页获取DockerCluster列表
export const getDockerClusterList = (params) => {
  return service({
    url: '/dockerCluster/getDockerClusterList',
    method: 'get',
    params
  })
}

// @Tags DockerCluster
// @Summary 获取Docker集群凭证
export const getDockerClusterCredentials = (params) => {
  return service({
    url: '/dockerCluster/getDockerClusterCredentials',
    method: 'get',
    params
  })
}

// @Tags DockerCluster
// @Summary 获取所有Docker集群（用于下拉选择）
export const getAllDockerClusters = () => {
  return service({
    url: '/dockerCluster/getAllDockerClusters',
    method: 'get'
  })
}
