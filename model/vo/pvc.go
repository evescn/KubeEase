package vo

import (
	"KubeEase/model/bo"
	corev1 "k8s.io/api/core/v1"
)

type PvcListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	bo.Common
}

type PvcRequest struct {
	PvcName string `json:"pvc_name" form:"pvc_name" binding:"required"`
	bo.Common
}

type PvcUpdateRequest struct {
	Content string `json:"content" form:"content" binding:"required"`
	bo.Common
}

// PvcsResp 定义列表的返回类型
type PvcsResp struct {
	Items []corev1.PersistentVolumeClaim `json:"items"`
	Total int                            `json:"total"`
}
