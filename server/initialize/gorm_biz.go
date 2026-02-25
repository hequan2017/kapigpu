package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/docker"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(docker.DockerCluster{}, docker.ComputeNode{})
	if err != nil {
		return err
	}
	return nil
}
