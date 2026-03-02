package config

// DockerCluster Docker集群配置
type DockerCluster struct {
	EncryptKey string `mapstructure:"encrypt-key" json:"encrypt-key" yaml:"encrypt-key"` // 加密密钥（16/24/32字节）
}
