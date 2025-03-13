package vo

import (
	"KubeEase/model/bo"
	corev1 "k8s.io/api/core/v1"
)

type PodListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	*bo.Common
}

type PodRequest struct {
	PodName string `json:"pod_name" form:"pod_name" binding:"required"`
	*bo.Common
}

type PodUpdateRequest struct {
	Content string `json:"content" binding:"required"`
	*bo.Common
}

type PodLogRequest struct {
	ContainerName string `json:"container_name" form:"container_name" binding:"required"`
	*PodRequest
}

type PodsResp struct {
	Items []corev1.Pod `json:"items"`
	Total int          `json:"total"`
}
