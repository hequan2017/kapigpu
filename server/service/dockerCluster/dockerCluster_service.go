package dockerCluster

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dockerCluster"
	dockerClusterReq "github.com/flipped-aurora/gin-vue-admin/server/model/dockerCluster/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/crypto"
	"go.uber.org/zap"
)

type DockerClusterService struct {
}

// getEncryptKey 获取加密密钥
// 从配置文件读取，如果没有配置则使用默认密钥
// 必须是16、24或32字节长度
func getEncryptKey() string {
	key := global.GVA_CONFIG.DockerCluster.EncryptKey
	if key == "" {
		// 默认密钥（仅用于开发环境，生产环境必须在config.yaml中配置）
		key = "gin-vue-admin-k8s-cluster-key!"
	}
	return key
}

// encryptSensitiveFields 加密敏感字段
func encryptSensitiveFields(cluster *dockerCluster.DockerCluster) error {
	key := getEncryptKey()

	// 加密CA证书
	if cluster.CaCert != "" {
		encrypted, err := crypto.AesEncrypt(cluster.CaCert, key)
		if err != nil {
			return err
		}
		cluster.CaCert = encrypted
	}

	// 加密客户端证书
	if cluster.ClientCert != "" {
		encrypted, err := crypto.AesEncrypt(cluster.ClientCert, key)
		if err != nil {
			return err
		}
		cluster.ClientCert = encrypted
	}

	// 加密客户端私钥
	if cluster.ClientKey != "" {
		encrypted, err := crypto.AesEncrypt(cluster.ClientKey, key)
		if err != nil {
			return err
		}
		cluster.ClientKey = encrypted
	}

	return nil
}

// decryptSensitiveFields 解密敏感字段
func decryptSensitiveFields(cluster *dockerCluster.DockerCluster) error {
	key := getEncryptKey()

	// 解密CA证书
	if cluster.CaCert != "" {
		decrypted, err := crypto.AesDecrypt(cluster.CaCert, key)
		if err != nil {
			return err
		}
		cluster.CaCert = decrypted
	}

	// 解密客户端证书
	if cluster.ClientCert != "" {
		decrypted, err := crypto.AesDecrypt(cluster.ClientCert, key)
		if err != nil {
			return err
		}
		cluster.ClientCert = decrypted
	}

	// 解密客户端私钥
	if cluster.ClientKey != "" {
		decrypted, err := crypto.AesDecrypt(cluster.ClientKey, key)
		if err != nil {
			return err
		}
		cluster.ClientKey = decrypted
	}

	return nil
}

// CreateDockerCluster 创建Docker集群记录
func (dockerClusterService *DockerClusterService) CreateDockerCluster(dc *dockerCluster.DockerCluster) (err error) {
	// 加密敏感字段
	if err = encryptSensitiveFields(dc); err != nil {
		global.GVA_LOG.Error("加密敏感字段失败", zap.Error(err))
		return errors.New("加密敏感字段失败: " + err.Error())
	}

	// 检查集群名称是否已存在
	var count int64
	if err = global.GVA_DB.Model(&dockerCluster.DockerCluster{}).Where("cluster_name = ?", dc.ClusterName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("集群名称已存在")
	}

	return global.GVA_DB.Create(dc).Error
}

// DeleteDockerCluster 删除Docker集群记录
func (dockerClusterService *DockerClusterService) DeleteDockerCluster(dc dockerCluster.DockerCluster) (err error) {
	return global.GVA_DB.Delete(&dc).Error
}

// DeleteDockerClusterByIds 批量删除Docker集群记录
func (dockerClusterService *DockerClusterService) DeleteDockerClusterByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Delete(&[]dockerCluster.DockerCluster{}, "id in ?", ids.Ids).Error
}

// UpdateDockerCluster 更新Docker集群记录
func (dockerClusterService *DockerClusterService) UpdateDockerCluster(dc dockerCluster.DockerCluster) (err error) {
	// 获取原有记录，判断是否更新了敏感字段
	var oldCluster dockerCluster.DockerCluster
	if err = global.GVA_DB.First(&oldCluster, dc.ID).Error; err != nil {
		return err
	}

	// 如果CA证书不为空且与原有记录不同（不是加密后的），则重新加密
	// 简单的判断方式：如果长度小于200，认为是原始证书需要加密
	if dc.CaCert != "" && len(dc.CaCert) < 200 {
		if err = encryptSensitiveFields(&dc); err != nil {
			global.GVA_LOG.Error("加密CA证书失败", zap.Error(err))
			return errors.New("加密CA证书失败: " + err.Error())
		}
	} else if dc.CaCert == "" {
		// 如果传空，保留原值
		dc.CaCert = oldCluster.CaCert
	}

	key := getEncryptKey()

	// 客户端证书
	if dc.ClientCert != "" && len(dc.ClientCert) < 200 {
		encrypted, err := crypto.AesEncrypt(dc.ClientCert, key)
		if err != nil {
			global.GVA_LOG.Error("加密客户端证书失败", zap.Error(err))
			return errors.New("加密客户端证书失败: " + err.Error())
		}
		dc.ClientCert = encrypted
	} else if dc.ClientCert == "" {
		dc.ClientCert = oldCluster.ClientCert
	}

	// 客户端私钥
	if dc.ClientKey != "" && len(dc.ClientKey) < 200 {
		encrypted, err := crypto.AesEncrypt(dc.ClientKey, key)
		if err != nil {
			global.GVA_LOG.Error("加密客户端私钥失败", zap.Error(err))
			return errors.New("加密客户端私钥失败: " + err.Error())
		}
		dc.ClientKey = encrypted
	} else if dc.ClientKey == "" {
		dc.ClientKey = oldCluster.ClientKey
	}

	// 更新非敏感字段
	updates := map[string]interface{}{
		"cluster_name": dc.ClusterName,
		"remark":       dc.Remark,
		"ca_cert":      dc.CaCert,
		"client_cert":  dc.ClientCert,
		"client_key":   dc.ClientKey,
	}

	return global.GVA_DB.Model(&dc).Updates(updates).Error
}

// GetDockerCluster 根据ID获取Docker集群记录
func (dockerClusterService *DockerClusterService) GetDockerCluster(id uint) (dc dockerCluster.DockerCluster, err error) {
	err = global.GVA_DB.First(&dc, id).Error
	if err != nil {
		return dc, err
	}

	// 解密敏感字段
	if err = decryptSensitiveFields(&dc); err != nil {
		global.GVA_LOG.Error("解密敏感字段失败", zap.Error(err))
		// 解密失败不返回错误，记录日志即可（可能是数据本身就不是加密的）
	}

	return dc, nil
}

// GetDockerClusterInfoList 分页获取Docker集群记录
func (dockerClusterService *DockerClusterService) GetDockerClusterInfoList(info dockerClusterReq.DockerClusterSearch) (list []dockerCluster.DockerCluster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&dockerCluster.DockerCluster{})
	var dcList []dockerCluster.DockerCluster

	// 如果有条件搜索
	if info.ClusterName != "" {
		db = db.Where("cluster_name LIKE ?", "%"+info.ClusterName+"%")
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 这里不返回解密后的敏感信息，只返回集群名称和备注等基本信息
	err = db.Limit(limit).Offset(offset).Find(&dcList).Error

	// 清空敏感字段（列表查询不返回证书信息）
	for i := range dcList {
		dcList[i].CaCert = ""
		dcList[i].ClientCert = ""
		dcList[i].ClientKey = ""
	}

	return dcList, total, err
}

// GetDockerClusterCredentials 获取集群凭证（用于连接Docker集群）
func (dockerClusterService *DockerClusterService) GetDockerClusterCredentials(id uint) (dc dockerCluster.DockerCluster, err error) {
	err = global.GVA_DB.First(&dc, id).Error
	if err != nil {
		return dc, err
	}

	// 解密敏感字段
	if err = decryptSensitiveFields(&dc); err != nil {
		global.GVA_LOG.Error("解密敏感字段失败", zap.Error(err))
		return dc, errors.New("解密凭证失败: " + err.Error())
	}

	return dc, nil
}

// GetAllDockerClusters 获取所有集群（用于下拉选择）
func (dockerClusterService *DockerClusterService) GetAllDockerClusters() (list []dockerCluster.DockerCluster, err error) {
	err = global.GVA_DB.Model(&dockerCluster.DockerCluster{}).
		Select("id, cluster_name, remark, created_at, updated_at").
		Find(&list).Error
	return list, err
}
