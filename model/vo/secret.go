package vo

import (
	"KubeEase/model/bo"
	corev1 "k8s.io/api/core/v1"
)

type SecretListRequest struct {
	FilterName string `form:"filter_name"`
	Size       int    `form:"size" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	bo.Common
}

type SecretRequest struct {
	SecretName string `json:"secret_name" form:"secret_name" binding:"required"`
	bo.Common
}

type SecretUpdateRequest struct {
	Content string `json:"content" form:"content" binding:"required"`
	bo.Common
}

// SecretsResp 定义列表的返回类型
type SecretsResp struct {
	Items []corev1.Secret `json:"items"`
	Total int             `json:"total"`
}
