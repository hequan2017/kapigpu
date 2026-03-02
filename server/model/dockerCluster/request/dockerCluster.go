package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// DockerClusterSearch 分页条件查询请求
// 支持按集群名称和备注进行模糊搜索
type DockerClusterSearch struct {
	ClusterName string `json:"clusterName" form:"clusterName"` // 集群名字（模糊搜索）
	Remark      string `json:"remark" form:"remark"`             // 备注（模糊搜索）
	request.PageInfo
}
