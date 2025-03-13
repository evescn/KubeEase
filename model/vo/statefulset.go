package vo

import (
	"KubeEase/model/bo"
	appsv1 "k8s.io/api/apps/v1"
)

type StatefulSetListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	*bo.Common
}

type StatefulSetRequest struct {
	StsName string `json:"sts_name" form:"sts_name" binding:"required"`
	*bo.Common
}

type StatefulSetUpdateRequest struct {
	Content string `json:"content" binding:"required"`
	*bo.Common
}

// StatefulSetResp 定义列表的返回类型
type StatefulSetResp struct {
	Items []appsv1.StatefulSet `json:"items"`
	Total int                  `json:"total"`
}
