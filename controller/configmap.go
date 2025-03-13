package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var ConfigMap configmap

type configmap struct{}

// GetConfigMaps 获取 ConfigMap 列表
func (p *configmap) GetConfigMaps(c *gin.Context) {
	params := &vo.ConfigMapListRequest{}

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
	data, err := service.ConfigMap.GetConfigMaps(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 ConfigMap 列表成功")
}

// GetConfigMapDetail 获取 ConfigMap 详情
func (p *configmap) GetConfigMapDetail(c *gin.Context) {
	params := &vo.ConfigMapRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.ConfigMap.GetConfigMapDetail(client, params.CmName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 ConfigMap 详情成功")
}

// DeleteConfigMap 删除 ConfigMap
func (p *configmap) DeleteConfigMap(c *gin.Context) {
	params := &vo.ConfigMapRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.ConfigMap.DeleteConfigMap(client, params.CmName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 ConfigMap 成功")
}

// UpdateConfigMap 更新 ConfigMap
func (p *configmap) UpdateConfigMap(c *gin.Context) {
	params := &vo.ConfigMapUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.ConfigMap.UpdateConfigMap(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 ConfigMap 成功")
}
