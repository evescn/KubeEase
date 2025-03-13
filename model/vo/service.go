package vo

import (
	"KubeEase/model/bo"
	corev1 "k8s.io/api/core/v1"
)

type ServiceListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	*bo.Common
}

type ServiceRequest struct {
	ServiceName string `json:"svc_name" form:"svc_name" binding:"required"`
	*bo.Common
}

type ServiceUpdateRequest struct {
	Content string `json:"content" binding:"required"`
	*bo.Common
}

// ServicesResp 定义列表的返回类型
type ServicesResp struct {
	Items []corev1.Service `json:"items"`
	Total int              `json:"total"`
}

type ServiceCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Label         map[string]string `json:"label"`
	Cluster       string            `json:"cluster"`
}
