package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Deployment deployment

type deployment struct{}

// GetDeployments 获取 Deployment 列表
func (p *deployment) GetDeployments(c *gin.Context) {
	params := &vo.DeploymentListRequest{}

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
	data, err := service.Deployment.GetDeployments(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Deployment 列表成功")
}

// GetDeploymentDetail 获取 Deployment 详情
func (p *deployment) GetDeploymentDetail(c *gin.Context) {
	params := &vo.DeploymentRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Deployment.GetDeploymentDetail(client, params.DeploymentName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Deployment 详情成功")
}

// DeleteDeployment 删除 Deployment
func (p *deployment) DeleteDeployment(c *gin.Context) {
	params := &vo.DeploymentRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Deployment.DeleteDeployment(client, params.DeploymentName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 Deployment 成功")
}

// UpdateDeployment 更新 Deployment
func (p *deployment) UpdateDeployment(c *gin.Context) {
	params := &vo.DeploymentUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Deployment.UpdateDeployment(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 Deployment 成功")
}

// ScaleDeployment 修改 Deployment 副本数
func (p *deployment) ScaleDeployment(c *gin.Context) {
	params := &vo.DeploymentScaleRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.Deployment.ScaleDeployment(client, params.ScaleNum, params.DeploymentName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "修改 Deployment 副本数成功")
}

// RestartDeployment 重启 Deployment
func (p *deployment) RestartDeployment(c *gin.Context) {
	params := &vo.DeploymentRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Deployment.RestartDeployment(client, params.DeploymentName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "重启 Deployment 成功")
}

// CreateDeployment 创建 Deployment
func (p *deployment) CreateDeployment(c *gin.Context) {
	params := &vo.DeploymentCreate{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.Deployment.CreateDeployment(client, params)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "创建 Deployment 成功")
}
