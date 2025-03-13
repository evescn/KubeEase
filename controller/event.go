package controller

import (
	"KubeEase/common"
	"KubeEase/model/vo"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var Event event

type event struct{}

// GetList 获取 event 列表，支持过滤、排序、分页
func (*event) GetList(c *gin.Context) {
	params := &vo.EventListRequest{}
	if err := c.Bind(params); err != nil {
		common.ResponseParamInvalid(c, err)
		return
	}
	data, err := service.Event.GetList(params.Name, params.Cluster, params.Page, params.Size)
	if err != nil {
		common.ResponseFailed(c, err)
		return
	}
	common.ResponseOk(c, data, "获取 Event 列表成功")
}
