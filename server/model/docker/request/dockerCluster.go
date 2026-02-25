
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DockerClusterSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      ClusterName  *string `json:"clusterName" form:"clusterName"` 
    request.PageInfo
}
