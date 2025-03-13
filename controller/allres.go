package controller

import (
	"KubeEase/common"
	"KubeEase/service"
	"github.com/gin-gonic/gin"
)

var AllRes allRes

type allRes struct{}

func (*allRes) GetAllNum(ctx *gin.Context) {
	params := new(struct {
		Cluster string `form:"cluster"`
	})
	if err := ctx.Bind(params); err != nil {
		common.ResponseParamInvalid(ctx, err)
		return
	}
	client, err := service.K8s.GetClient(params.Cluster)
	if err != nil {
		common.ResponseFailed(ctx, err)
		return
	}
	data, errs := service.AllRes.GetAllNum(client)
	if len(errs) > 0 {
		common.ResponseFailed(ctx, err)
		return
	}
	common.ResponseOk(ctx, data, "获取资源数量成功")
}
