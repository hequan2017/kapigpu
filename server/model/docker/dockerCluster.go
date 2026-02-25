
// 自动生成模板DockerCluster
package docker
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Docker集群管理 结构体  DockerCluster
type DockerCluster struct {
    global.GVA_MODEL
  ClusterName  *string `json:"clusterName" form:"clusterName" gorm:"comment:集群名字;column:cluster_name;size:128;" binding:"required"`  //集群名字
  CaCert  *string `json:"caCert" form:"caCert" gorm:"comment:CA证书;column:ca_cert;size:8192;"`  //CA证书
  ClientCert  *string `json:"clientCert" form:"clientCert" gorm:"comment:客户端证书;column:client_cert;size:8192;"`  //客户端证书
  ClientKey  *string `json:"clientKey" form:"clientKey" gorm:"comment:客户端私钥;column:client_key;size:8192;"`  //客户端私钥
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:512;"`  //备注
}


// TableName Docker集群管理 DockerCluster自定义表名 docker_cluster
func (DockerCluster) TableName() string {
    return "docker_cluster"
}





