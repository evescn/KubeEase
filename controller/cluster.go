package controller

import (
	"KubeEase/common"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
	"sort"
)

var Cluster cluster

type cluster struct{}

func (*cluster) GetClusters(c *gin.Context) {
	list := make([]string, 0)
	for key := range service.K8s.ClientMap {
		list = append(list, key)
	}
	sort.Strings(list)
	common.ResponseOk(c, list, "获取集群信息成功")
}
