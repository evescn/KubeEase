package vo

import (
	"KubeEase/model/bo"
	appsv1 "k8s.io/api/apps/v1"
)

type DeploymentListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	*bo.Common
}

type DeploymentRequest struct {
	DeploymentName string `json:"deployment_name" form:"deployment_name" binding:"required"`
	*bo.Common
}

type DeploymentUpdateRequest struct {
	Content string `json:"content" binding:"required"`
	*bo.Common
}

type DeploymentScaleRequest struct {
	ScaleNum int `json:"scale_num" form:"scale_num" binding:"required"`
	*DeploymentRequest
}

type DeploymentResp struct {
	Items []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

// DeploymentCreate 定义 Deployment 创建的结构体
type DeploymentCreate struct {
	Name          string            `json:"name"`
	Replicas      int32             `json:"replicas"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	Cpu           string            `json:"cpu"`
	Memory        string            `json:"memory"`
	ContainerPort int32             `json:"container_port"`
	HealthCheck   bool              `json:"health_check"`
	HealthPath    string            `json:"health_path"`
	*bo.Common
}
