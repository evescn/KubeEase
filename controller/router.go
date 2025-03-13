package controller

import (
	"KubeEase/logger"
	"KubeEase/middleware"
	"KubeEase/settings"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router router

type router struct{}

func (*router) Setup() *gin.Engine {
	// 初始化gin对象
	if settings.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 修改日志格式
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 跨域中间件
	r.Use(middleware.Cors())
	// JWT登陆验证中间件
	//r.Use(middle.JWTAuth())

	r.GET("/testApi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "testApi success!",
			"data": nil,
		})
	})

	r.GET("/version", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, settings.Conf.Version)
	})

	k8s := r.Group("/api/k8s/")
	//集群
	k8s.GET("/clusters", Cluster.GetClusters)

	// AllRes 操作
	k8s.GET("/allres", AllRes.GetAllNum)

	// Event 操作
	k8s.GET("/events", Event.GetList)

	// Node 操作
	k8s.GET("/nodes", Node.GetNodes)
	k8s.GET("/node/detail", Node.GetNodeDetail)

	// Namespace 操作
	k8s.GET("/namespaces", Namespace.GetNamespaces)
	k8s.GET("/namespace/detail", Namespace.GetNamespaceDetail)
	k8s.DELETE("/namespace/del", Namespace.DeleteNamespace)

	// PV 操作
	k8s.GET("/pvs", Pv.GetPvs)
	k8s.GET("/pv/detail", Pv.GetPvDetail)
	k8s.DELETE("/pv/del", Pv.DeletePv)

	// Pod 操作
	k8s.GET("/pods", Pod.GetPods)
	k8s.GET("/pod/detail", Pod.GetPodDetail)
	k8s.PUT("/pod/update", Pod.UpdatePod)
	k8s.DELETE("/pod/del", Pod.DeletePod)
	k8s.GET("/pod/log", Pod.GetPodLog)
	k8s.GET("/pod/container", Pod.GetPodContainer)

	// Deployment 操作
	k8s.GET("/deployments", Deployment.GetDeployments)
	k8s.GET("/deployment/detail", Deployment.GetDeploymentDetail)
	k8s.DELETE("/deployment/del", Deployment.DeleteDeployment)
	k8s.PUT("/deployment/update", Deployment.UpdateDeployment)
	k8s.PUT("/deployment/scale", Deployment.ScaleDeployment)
	k8s.PUT("/deployment/restart", Deployment.RestartDeployment)
	k8s.POST("/deployment/create", Deployment.CreateDeployment)

	// DaemonSet 操作
	k8s.GET("/daemonsets", DaemonSet.GetDaemonSets)
	k8s.GET("/daemonset/detail", DaemonSet.GetDaemonSetDetail)
	k8s.DELETE("/daemonset/del", DaemonSet.DeleteDaemonSet)
	k8s.PUT("/daemonset/update", DaemonSet.UpdateDaemonSet)

	// StatefulSet 操作
	k8s.GET("/statefulsets", StatefulSet.GetStatefulSets)
	k8s.GET("/statefulset/detail", StatefulSet.GetStatefulSetDetail)
	k8s.DELETE("/statefulset/del", StatefulSet.DeleteStatefulSet)
	k8s.PUT("/statefulset/update", StatefulSet.UpdateStatefulSet)

	// Service 操作
	k8s.GET("/services", Servicev1.GetServices)
	k8s.GET("/service/detail", Servicev1.GetServiceDetail)
	k8s.DELETE("/service/del", Servicev1.DeleteService)
	k8s.PUT("/service/update", Servicev1.UpdateService)
	k8s.POST("/service/create", Servicev1.CreateService)

	// Ingress 操作
	k8s.GET("/ingresses", Ingress.GetIngresses)
	k8s.GET("/ingress/detail", Ingress.GetIngressDetail)
	k8s.DELETE("/ingress/del", Ingress.DeleteIngress)
	k8s.PUT("/ingress/update", Ingress.UpdateIngress)
	k8s.POST("/ingress/create", Ingress.CreateIngress)

	// ConfigMap 操作
	k8s.GET("/configmaps", ConfigMap.GetConfigMaps)
	k8s.GET("/configmap/detail", ConfigMap.GetConfigMapDetail)
	k8s.DELETE("/configmap/del", ConfigMap.DeleteConfigMap)
	k8s.PUT("/configmap/update", ConfigMap.UpdateConfigMap)

	// Secret 操作
	k8s.GET("/secrets", Secret.GetSecrets)
	k8s.GET("/secret/detail", Secret.GetSecretDetail)
	k8s.DELETE("/secret/del", Secret.DeleteSecret)
	k8s.PUT("/secret/update", Secret.UpdateSecret)

	// PVC 操作
	k8s.GET("/pvcs", Pvc.GetPvcs)
	k8s.GET("/pvc/detail", Pvc.GetPvcDetail)
	k8s.DELETE("/pvc/del", Pvc.DeletePvc)
	k8s.PUT("/pvc/update", Pvc.UpdatePvc)

	//// Helm 应用商店
	//GET("/api/helmstore/releases", controller.HelmStore.ListReleases).
	//GET("/api/helmstore/release/detail", controller.HelmStore.DetailRelease).
	//POST("/api/helmstore/release/install", controller.HelmStore.InstallRelease).
	//DELETE("/api/helmstore/release/uninstall", controller.HelmStore.UninstallRelease).
	//GET("/api/helmstore/charts", controller.HelmStore.ListCharts).
	//POST("/api/helmstore/chart/add", controller.HelmStore.AddChart).
	//PUT("/api/helmstore/chart/update", controller.HelmStore.UpdateChart).
	//DELETE("/api/helmstore/chart/del", controller.HelmStore.DeleteChart).
	//POST("/api/helmstore/chartfile/upload", controller.HelmStore.UploadChartFile).
	//DELETE("/api/helmstore/chartfile/del", controller.HelmStore.DeleteChartFile)

	////登录验证，路由权限信息
	//POST("/api/login", controller.Login.Auth).

	return r
}
