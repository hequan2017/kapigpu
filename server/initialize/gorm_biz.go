package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dockerCluster"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(
		dockerCluster.DockerCluster{},
	)
	if err != nil {
		return err
	}
	return nil
}
