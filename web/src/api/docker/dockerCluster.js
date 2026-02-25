import service from '@/utils/request'
// @Tags DockerCluster
// @Summary 创建Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DockerCluster true "创建Docker集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dockerCluster/createDockerCluster [post]
export const createDockerCluster = (data) => {
  return service({
    url: '/dockerCluster/createDockerCluster',
    method: 'post',
    data
  })
}

// @Tags DockerCluster
// @Summary 删除Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DockerCluster true "删除Docker集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dockerCluster/deleteDockerCluster [delete]
export const deleteDockerCluster = (params) => {
  return service({
    url: '/dockerCluster/deleteDockerCluster',
    method: 'delete',
    params
  })
}

// @Tags DockerCluster
// @Summary 批量删除Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Docker集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dockerCluster/deleteDockerCluster [delete]
export const deleteDockerClusterByIds = (params) => {
  return service({
    url: '/dockerCluster/deleteDockerClusterByIds',
    method: 'delete',
    params
  })
}

// @Tags DockerCluster
// @Summary 更新Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DockerCluster true "更新Docker集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dockerCluster/updateDockerCluster [put]
export const updateDockerCluster = (data) => {
  return service({
    url: '/dockerCluster/updateDockerCluster',
    method: 'put',
    data
  })
}

// @Tags DockerCluster
// @Summary 用id查询Docker集群管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DockerCluster true "用id查询Docker集群管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dockerCluster/findDockerCluster [get]
export const findDockerCluster = (params) => {
  return service({
    url: '/dockerCluster/findDockerCluster',
    method: 'get',
    params
  })
}

// @Tags DockerCluster
// @Summary 分页获取Docker集群管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Docker集群管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dockerCluster/getDockerClusterList [get]
export const getDockerClusterList = (params) => {
  return service({
    url: '/dockerCluster/getDockerClusterList',
    method: 'get',
    params
  })
}

// @Tags DockerCluster
// @Summary 不需要鉴权的Docker集群管理接口
// @Accept application/json
// @Produce application/json
// @Param data query dockerReq.DockerClusterSearch true "分页获取Docker集群管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dockerCluster/getDockerClusterPublic [get]
export const getDockerClusterPublic = () => {
  return service({
    url: '/dockerCluster/getDockerClusterPublic',
    method: 'get',
  })
}
