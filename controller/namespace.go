package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Namespace namespace

type namespace struct{}

// GetNamespaces 获取 Namespace 列表
func (n *namespace) GetNamespaces(c *gin.Context) {
	params := &vo.NamespaceListRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Namespace.GetNamespaces(client, params.FilterName, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Namespace 列表成功")
}

// GetNamespaceDetail 获取 Namespace 列表
func (n *namespace) GetNamespaceDetail(c *gin.Context) {
	params := &vo.NamespaceRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Namespace.GetNamespaceDetail(client, params.NamespaceName)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Namespace 详情成功")
}

// DeleteNamespace 删除 Namespace
func (n *namespace) DeleteNamespace(c *gin.Context) {
	params := &vo.NamespaceRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Namespace.DeleteNamespace(client, params.NamespaceName)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 Namespace 成功")
}
