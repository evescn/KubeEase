package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Servicev1 servicev1

type servicev1 struct{}

// GetServices 获取 Service 列表
func (s *servicev1) GetServices(c *gin.Context) {
	params := &vo.ServiceListRequest{}

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
	data, err := service.Servicev1.GetServices(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Service 列表成功")
}

// GetServiceDetail 获取 Service 详情
func (s *servicev1) GetServiceDetail(c *gin.Context) {
	params := &vo.ServiceRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Servicev1.GetServiceDetail(client, params.ServiceName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Service 详情成功")
}

// DeleteService 删除 Service
func (s *servicev1) DeleteService(c *gin.Context) {
	params := &vo.ServiceRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Servicev1.DeleteService(client, params.ServiceName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 Service 成功")
}

// UpdateService 更新 Service
func (s *servicev1) UpdateService(c *gin.Context) {
	params := &vo.ServiceUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error(fmt.Sprintf("绑定参数失败， %v\n", err))
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Servicev1.UpdateService(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 Service 成功")
}

// CreateService 创建 Service
func (s *servicev1) CreateService(c *gin.Context) {
	params := &vo.ServiceCreate{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Servicev1.CreateService(client, params)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "创建 Service 成功")
}
