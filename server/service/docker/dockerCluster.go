
package docker

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/docker"
	dockerReq "github.com/flipped-aurora/gin-vue-admin/server/model/docker/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DockerClusterService struct{}

func getEncryptKey() string {
	return global.GVA_CONFIG.JWT.SigningKey
}

func encryptField(val *string) {
	if val == nil || *val == "" {
		return
	}
	enc, err := utils.AesEncrypt(*val, getEncryptKey())
	if err == nil {
		*val = enc
	}
}

func decryptField(val *string) {
	if val == nil || *val == "" {
		return
	}
	dec, err := utils.AesDecrypt(*val, getEncryptKey())
	if err == nil {
		*val = dec
	}
}

func encryptCluster(c *docker.DockerCluster) {
	encryptField(c.CaCert)
	encryptField(c.ClientCert)
	encryptField(c.ClientKey)
}

func decryptCluster(c *docker.DockerCluster) {
	decryptField(c.CaCert)
	decryptField(c.ClientCert)
	decryptField(c.ClientKey)
}

// CreateDockerCluster 创建Docker集群管理记录
func (dockerClusterService *DockerClusterService) CreateDockerCluster(ctx context.Context, dockerCluster *docker.DockerCluster) (err error) {
	encryptCluster(dockerCluster)
	err = global.GVA_DB.Create(dockerCluster).Error
	return err
}

// DeleteDockerCluster 删除Docker集群管理记录
func (dockerClusterService *DockerClusterService) DeleteDockerCluster(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&docker.DockerCluster{}, "id = ?", ID).Error
	return err
}

// DeleteDockerClusterByIds 批量删除Docker集群管理记录
func (dockerClusterService *DockerClusterService) DeleteDockerClusterByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]docker.DockerCluster{}, "id in ?", IDs).Error
	return err
}

// UpdateDockerCluster 更新Docker集群管理记录
func (dockerClusterService *DockerClusterService) UpdateDockerCluster(ctx context.Context, dockerCluster docker.DockerCluster) (err error) {
	encryptCluster(&dockerCluster)
	err = global.GVA_DB.Model(&docker.DockerCluster{}).Where("id = ?", dockerCluster.ID).Updates(&dockerCluster).Error
	return err
}

// GetDockerCluster 根据ID获取Docker集群管理记录
func (dockerClusterService *DockerClusterService) GetDockerCluster(ctx context.Context, ID string) (dockerCluster docker.DockerCluster, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dockerCluster).Error
	if err == nil {
		decryptCluster(&dockerCluster)
	}
	return
}
// GetDockerClusterInfoList 分页获取Docker集群管理记录
// Author [yourname](https://github.com/yourname)
func (dockerClusterService *DockerClusterService)GetDockerClusterInfoList(ctx context.Context, info dockerReq.DockerClusterSearch) (list []docker.DockerCluster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&docker.DockerCluster{})
    var dockerClusters []docker.DockerCluster
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.ClusterName != nil && *info.ClusterName != "" {
        db = db.Where("cluster_name LIKE ?", "%"+ *info.ClusterName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&dockerClusters).Error
	return  dockerClusters, total, err
}
func (dockerClusterService *DockerClusterService)GetDockerClusterPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
