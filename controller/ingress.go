package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Ingress ingress

type ingress struct{}

// GetIngresses 获取 Ingress 列表
func (i *ingress) GetIngresses(c *gin.Context) {
	params := &vo.IngressListRequest{}

	//绑定参数
	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	//获取client
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	//调用Ingress方法，获取列表
	data, err := service.Ingress.GetIngresses(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Ingress 列表成功")
}

// GetIngressDetail 获取 Ingress 详情
func (i *ingress) GetIngressDetail(c *gin.Context) {
	params := &vo.IngressRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Ingress.GetIngressDetail(client, params.IngressName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Ingress 详情成功")
}

// DeleteIngress 删除 Ingress
func (i *ingress) DeleteIngress(c *gin.Context) {
	params := &vo.IngressRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Ingress.DeleteIngress(client, params.IngressName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 Ingress 成功")
}

// UpdateIngress 更新 Ingress
func (i *ingress) UpdateIngress(c *gin.Context) {
	params := &vo.IngressUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Ingress.UpdateIngress(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 Ingress 成功")
}

// CreateIngress 创建 Ingress
func (i *ingress) CreateIngress(c *gin.Context) {
	params := &vo.IngressCreate{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Ingress.CreateIngress(client, params)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "创建 Ingress 成功")
}
