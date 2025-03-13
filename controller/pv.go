package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Pv pv

type pv struct{}

// GetPvs 获取 Pv 列表
func (p *pv) GetPvs(c *gin.Context) {
	params := &vo.PvListRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pv.GetPvs(client, params.FilterName, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 PV 列表成功")
}

// GetPvDetail 获取 Pv 列表
func (p *pv) GetPvDetail(c *gin.Context) {
	params := &vo.PvRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pv.GetPvDetail(client, params.PvName)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 PV 详情成功")
}

// DeletePv 删除 Pv
func (p *pv) DeletePv(c *gin.Context) {
	params := &vo.PvRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Pv.DeletePv(client, params.PvName)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 PV 成功")
}
