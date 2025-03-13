package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var DaemonSet daemonSet

type daemonSet struct{}

// GetDaemonSets 获取 DaemonSet 列表
func (d *daemonSet) GetDaemonSets(c *gin.Context) {
	params := &vo.DaemonSetListRequest{}

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

	//调用service方法，获取列表
	data, err := service.DaemonSet.GetDaemonSets(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 DaemonSet 列表成功")
}

// GetDaemonSetDetail 获取 DaemonSet 详情
func (d *daemonSet) GetDaemonSetDetail(c *gin.Context) {
	params := &vo.DaemonSetRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.DaemonSet.GetDaemonSetDetail(client, params.DsName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 DaemonSet 详情成功")
}

// DeleteDaemonSet 删除 DaemonSet
func (d *daemonSet) DeleteDaemonSet(c *gin.Context) {
	params := &vo.DaemonSetRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.DaemonSet.DeleteDaemonSet(client, params.DsName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 DaemonSet 成功")
}

// UpdateDaemonSet 更新 DaemonSet
func (d *daemonSet) UpdateDaemonSet(c *gin.Context) {
	params := &vo.DaemonSetUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.DaemonSet.UpdateDaemonSet(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 DaemonSet 成功")
}
