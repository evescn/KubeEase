package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Pvc pvc

type pvc struct{}

// GetPvcs 获取 Pvc 列表
func (p *pvc) GetPvcs(c *gin.Context) {
	params := &vo.PvcListRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pvc.GetPvcs(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Pvc 列表成功")
}

// GetPvcDetail 获取 Pvc 详情
func (p *pvc) GetPvcDetail(c *gin.Context) {
	params := &vo.PvcRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pvc.GetPvcDetail(client, params.PvcName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Pvc 详情成功")
}

// DeletePvc 删除 Pvc
func (p *pvc) DeletePvc(c *gin.Context) {
	params := &vo.PvcRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Pvc.DeletePvc(client, params.PvcName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 Pvc 成功")
}

// UpdatePvc 更新 Pvc
func (p *pvc) UpdatePvc(c *gin.Context) {
	params := &vo.PvcUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Pvc.UpdatePvc(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 Pvc 成功")
}
