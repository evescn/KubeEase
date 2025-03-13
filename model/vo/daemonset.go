package vo

import (
	"KubeEase/model/bo"
	appsv1 "k8s.io/api/apps/v1"
)

type DaemonSetListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	*bo.Common
}

type DaemonSetRequest struct {
	DsName string `json:"ds_name" form:"ds_name" binding:"required"`
	*bo.Common
}

type DaemonSetUpdateRequest struct {
	Content string `json:"content" binding:"required"`
	*bo.Common
}

// DaemonSetsResp 定义列表的返回类型
type DaemonSetsResp struct {
	Items []appsv1.DaemonSet `json:"items"`
	Total int                `json:"total"`
}
