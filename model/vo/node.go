package vo

import (
	corev1 "k8s.io/api/core/v1"
)

type NodeListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	Cluster    string `json:"cluster" form:"cluster" binding:"required"`
}

type NodeRequest struct {
	NodeName string `json:"node_name" form:"node_name" binding:"required"`
	Cluster  string `json:"cluster" form:"cluster" binding:"required"`
}

// NodesResp 定义列表的返回类型
type NodesResp struct {
	Items []corev1.Node `json:"items"`
	Total int           `json:"total"`
}
