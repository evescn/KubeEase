package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Pod pod

type pod struct{}

// GetPods 获取pod列表
func (p *pod) GetPods(c *gin.Context) {
	params := &vo.PodListRequest{}

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
	data, err := service.Pod.GetPods(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Pod 列表成功")
}

// GetPodDetail 获取pod详情
func (p *pod) GetPodDetail(c *gin.Context) {
	params := &vo.PodRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pod.GetPodDetail(client, params.PodName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Pod 详情成功")
}

// DeletePod 删除pod
func (p *pod) DeletePod(c *gin.Context) {
	params := &vo.PodRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Pod.DeletePod(client, params.PodName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 Pod 成功")
}

// UpdatePod 更新pod
func (p *pod) UpdatePod(c *gin.Context) {
	params := &vo.PodUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Pod.UpdatePod(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 Pod 成功")
}

// GetPodContainer 获取 Pod 容器名
func (p *pod) GetPodContainer(c *gin.Context) {
	params := &vo.PodRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pod.GetPodContainer(client, params.PodName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Pod 容器名成功")
}

// GetPodLog 获取 Pod 日志
func (p *pod) GetPodLog(c *gin.Context) {
	params := &vo.PodLogRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Pod.GetPodLog(client, params.ContainerName, params.PodName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Pod 容器日志成功")
}
