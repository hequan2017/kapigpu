
// 自动生成模板ComputeNode
package docker
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 算力节点 结构体  ComputeNode
type ComputeNode struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:名字;column:name;size:128;" binding:"required"`  //名字
  ClusterId  *int64 `json:"clusterId" form:"clusterId" gorm:"comment:关联Docker集群;column:cluster_id;"`  //集群
  Region  *string `json:"region" form:"region" gorm:"comment:区域;column:region;size:64;"`  //区域
  Cpu  *string `json:"cpu" form:"cpu" gorm:"comment:CPU;column:cpu;size:64;"`  //CPU
  Memory  *string `json:"memory" form:"memory" gorm:"comment:内存;column:memory;size:64;"`  //内存
  SystemDisk  *string `json:"systemDisk" form:"systemDisk" gorm:"comment:系统盘容量;column:system_disk;size:64;"`  //系统盘容量
  DataDisk  *string `json:"dataDisk" form:"dataDisk" gorm:"comment:数据盘容量;column:data_disk;size:64;"`  //数据盘容量
  PublicIp  *string `json:"publicIp" form:"publicIp" gorm:"comment:IP地址公网;column:public_ip;size:64;" binding:"required"`  //IP地址公网
  PrivateIp  *string `json:"privateIp" form:"privateIp" gorm:"comment:IP地址内网;column:private_ip;size:64;" binding:"required"`  //IP地址内网
  SshPort  *int64 `json:"sshPort" form:"sshPort" gorm:"default:22;comment:SSH端口;column:ssh_port;" binding:"required"`  //SSH端口
  Username  *string `json:"username" form:"username" gorm:"comment:用户名;column:username;size:128;"`  //用户名
  Password  *string `json:"password" form:"password" gorm:"comment:密码;column:password;size:256;"`  //密码
  GpuName  *string `json:"gpuName" form:"gpuName" gorm:"comment:显卡名称;column:gpu_name;size:128;"`  //显卡名称
  GpuCount  *int64 `json:"gpuCount" form:"gpuCount" gorm:"comment:显卡数量;column:gpu_count;"`  //显卡数量
  DockerAddr  *string `json:"dockerAddr" form:"dockerAddr" gorm:"comment:Docker连接地址;column:docker_addr;size:256;"`  //Docker连接地址
  IsOnline  *bool `json:"isOnline" form:"isOnline" gorm:"default:true;comment:是否上架;column:is_online;" binding:"required"`  //是否上架
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:512;"`  //备注
}


// TableName 算力节点 ComputeNode自定义表名 compute_node
func (ComputeNode) TableName() string {
    return "compute_node"
}





