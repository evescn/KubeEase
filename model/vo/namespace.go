package vo

import corev1 "k8s.io/api/core/v1"

type NamespaceListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	Cluster    string `json:"cluster" form:"cluster" binding:"required"`
}

type NamespaceRequest struct {
	NamespaceName string `json:"namespace_name" form:"namespace_name" binding:"required"`
	Cluster       string `json:"cluster" form:"cluster" binding:"required"`
}

// NamespacesResp 定义列表的返回类型
type NamespacesResp struct {
	Items []corev1.Namespace `json:"items"`
	Total int                `json:"total"`
}
