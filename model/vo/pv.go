package vo

import corev1 "k8s.io/api/core/v1"

type PvListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	Cluster    string `json:"cluster" form:"cluster" binding:"required"`
}

type PvRequest struct {
	PvName  string `json:"pv_name" form:"pv_name" binding:"required"`
	Cluster string `json:"cluster" form:"cluster" binding:"required"`
}

// PvsResp 定义列表的返回类型
type PvsResp struct {
	Items []corev1.PersistentVolume `json:"items"`
	Total int                       `json:"total"`
}
