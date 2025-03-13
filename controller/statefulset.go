package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var StatefulSet statefulSet

type statefulSet struct{}

// GetStatefulSets 获取 StatefulSet 列表
func (s *statefulSet) GetStatefulSets(c *gin.Context) {
	params := &vo.StatefulSetListRequest{}

	//绑定参数
	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	//获取client
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	//调用service方法，获取列表
	data, err := service.StatefulSet.GetStatefulSets(client, params.FilterName, params.Namespace, params.Size, params.Page)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 StatefulSet 列表成功")
}

// GetStatefulSetDetail 获取 StatefulSet 详情
func (s *statefulSet) GetStatefulSetDetail(c *gin.Context) {
	params := &vo.StatefulSetRequest{}

	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	data, err := service.StatefulSet.GetStatefulSetDetail(client, params.StsName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 StatefulSet 详情成功")
}

// DeleteStatefulSet 删除 StatefulSet
func (s *statefulSet) DeleteStatefulSet(c *gin.Context) {
	params := &vo.StatefulSetRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.StatefulSet.DeleteStatefulSet(client, params.StsName, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "删除 StatefulSet 成功")
}

// UpdateStatefulSet 更新 StatefulSet
func (s *statefulSet) UpdateStatefulSet(c *gin.Context) {
	params := &vo.StatefulSetUpdateRequest{}

	if err := c.ShouldBindJSON(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}

	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}

	err = service.StatefulSet.UpdateStatefulSet(client, params.Content, params.Namespace)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, nil, "更新 StatefulSet 成功")
}
