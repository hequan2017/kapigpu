package dockerCluster

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DockerCluster Docker集群管理表
// 结构体字段说明：
// - ClusterName: 集群名字，用于标识不同的Docker集群
// - CaCert: CA证书，加密存储，使用TEXT类型避免MySQL行大小限制
// - ClientCert: 客户端证书，加密存储，使用TEXT类型避免MySQL行大小限制
// - ClientKey: 客户端私钥，加密存储，使用TEXT类型避免MySQL行大小限制
// - Remark: 备注信息
type DockerCluster struct {
	global.GVA_MODEL
	ClusterName string `json:"clusterName" form:"clusterName" gorm:"column:cluster_name;comment:集群名字;size:191;uniqueIndex;"` // 集群名字
	CaCert      string `json:"caCert" form:"caCert" gorm:"column:ca_cert;comment:CA证书(加密存储);type:text;"`                // CA证书(加密存储) - 使用TEXT避免行大小限制
	ClientCert  string `json:"clientCert" form:"clientCert" gorm:"column:client_cert;comment:客户端证书(加密存储);type:text;"`  // 客户端证书(加密存储) - 使用TEXT避免行大小限制
	ClientKey   string `json:"clientKey" form:"clientKey" gorm:"column:client_key;comment:客户端私钥(加密存储);type:text;"`   // 客户端私钥(加密存储) - 使用TEXT避免行大小限制
	Remark      string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`                            // 备注
}

// TableName 指定表名
func (DockerCluster) TableName() string {
	return "docker_clusters"
}
