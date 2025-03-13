package vo

import (
	"KubeEase/model/bo"
	nwv1 "k8s.io/api/networking/v1"
)

type IngressListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	*bo.Common
}

type IngressRequest struct {
	IngressName string `json:"ingress_name" form:"ingress_name" binding:"required"`
	*bo.Common
}

type IngressUpdateRequest struct {
	Content string `json:"content" binding:"required"`
	*bo.Common
}

// IngressesResp 定义列表的返回类型
type IngressesResp struct {
	Items []nwv1.Ingress `json:"items"`
	Total int            `json:"total"`
}

// IngressCreate 定义 IngressCreate 的结构体
type IngressCreate struct {
	Name  string                 `json:"name"`
	Label map[string]string      `json:"label"`
	Hosts map[string][]*HttpPath `json:"hosts"`
	*bo.Common
}

// HttpPath 定义 ingress 的 path 结构体
type HttpPath struct {
	Path        string        `json:"path"`
	PathType    nwv1.PathType `json:"path_type"`
	ServiceName string        `json:"service_name"`
	ServicePort int32         `json:"service_port"`
}
