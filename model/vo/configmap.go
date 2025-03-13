package vo

import (
	"KubeEase/model/bo"
	corev1 "k8s.io/api/core/v1"
)

type ConfigMapListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	bo.Common
}

type ConfigMapRequest struct {
	CmName string `json:"cm_name" form:"cm_name" binding:"required"`
	bo.Common
}

type ConfigMapUpdateRequest struct {
	Content string `json:"content" form:"content" binding:"required"`
	bo.Common
}

// ConfigMapsResp 定义列表的返回类型
type ConfigMapsResp struct {
	Items []corev1.ConfigMap `json:"items"`
	Total int                `json:"total"`
}
