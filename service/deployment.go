package service

import (
	"KubeEase/model/vo"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	//"strconv"

	//"strconv"
	"time"
)

var Deployment deployment

type deployment struct{}

// toCells 方法用于将 deployment 类型数组，转换成 DataCell 类型数组
func (d *deployment) toCells(std []appsv1.Deployment) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = deploymentCell(std[i])
	}
	return cells
}

// fromCells 方法用于将 DataCell 类型数组，转换成 deployment 类型数组
func (d *deployment) fromCells(cells []DataCell) []appsv1.Deployment {
	deployments := make([]appsv1.Deployment, len(cells))
	for i := range cells {
		deployments[i] = appsv1.Deployment(cells[i].(deploymentCell))
	}
	return deployments
}

// GetDeployments 获取deployment列表
func (d *deployment) GetDeployments(client *kubernetes.Clientset, filterName, namespace string, size, page int) (deploymentResp *vo.DeploymentResp, err error) {
	deploymentList, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("获取 Deployment 列表失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取 Deployment 列表失败, %v\n", err))
	}
	//实例化dataSelector对象
	selectableData := &dataSelector{
		GenericDataList: d.toCells(deploymentList.Items),
		dataSelectQuery: &DataSelectQuery{
			FilterQuery: &FilterQuery{Name: filterName},
			PaginateQuery: &PaginateQuery{
				Size: size,
				Page: page,
			},
		},
	}

	//先过滤，filtered中的数据才是总数据，data中的数据是排序分页后的数据，可能每次只有10行
	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)

	//在排序和分页
	data := filtered.Sort().Paginate()
	//将[]DataCell类型的pod列表转为v1.pod列表
	deployments := d.fromCells(data.GenericDataList)
	return &vo.DeploymentResp{
		Items: deployments,
		Total: total,
	}, nil
}

// GetDeploymentDetail 获取 deployment 详情
func (d *deployment) GetDeploymentDetail(client *kubernetes.Clientset, deploymentName, namespace string) (deployment *appsv1.Deployment, err error) {
	deploymentDetail, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("获取Deployment详情失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取Deployment详情失败, %v\n", err))
	}
	return deploymentDetail, nil
}

// UpdateDeployment 更新 Deployment
// content就是deployment的整个json体
func (d *deployment) UpdateDeployment(client *kubernetes.Clientset, content, namespace string) (err error) {
	//content转成deployment结构体
	var deploy = &appsv1.Deployment{}
	//反序列化成deployment对象
	err = json.Unmarshal([]byte(content), &deploy)
	if err != nil {
		zap.L().Error(fmt.Sprintf("反序列化失败, %v\n", err))
		return errors.New(fmt.Sprintf("反序列化失败, %v\n", err))
	}
	//更新deployment
	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("更新 Deployment 失败, %v\n", err))
		return errors.New(fmt.Sprintf("更新 Deployment 失败, %v\n", err))
	}
	return nil
}

// DeleteDeployment 删除 Deployment
func (d *deployment) DeleteDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	err = client.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("删除 Deployment 失败, %v\n", err))
		return errors.New(fmt.Sprintf("删除 Deployment 失败, %v\n", err))
	}
	return nil
}

// ScaleDeployment 修改 Deployment 副本数
func (d *deployment) ScaleDeployment(client *kubernetes.Clientset, scaleNum int, deploymentName, namespace string) (replica int32, err error) {
	scale, err := client.AppsV1().Deployments(namespace).GetScale(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("获取Deployment副本信息失败, %v\n", err))
		return 0, errors.New(fmt.Sprintf("获取Deployment副本信息失败, %v\n", err))
	}
	scale.Spec.Replicas = int32(scaleNum)
	updateScale, err := client.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), deploymentName, scale, metav1.UpdateOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("更新Deployment副本信息失败, %v\n", err))
		return 0, errors.New(fmt.Sprintf("更新Deployment副本信息失败, %v\n", err))
	}
	return updateScale.Spec.Replicas, nil
}

// RestartDeployment 重启 Deployment
func (d *deployment) RestartDeployment(client *kubernetes.Clientset, deploymentName, namespace string) (err error) {
	//构造 patch 数据，更新 pod 模板 annotations
	now := time.Now().Format(time.RFC3339)
	patchData := fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, now)

	//调用patch方法更新deployment
	_, err = client.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName,
		types.StrategicMergePatchType, []byte(patchData), metav1.PatchOptions{})

	if err != nil {
		zap.L().Error(fmt.Sprintf("重启Deployment失败, %v\n", err))
		return errors.New(fmt.Sprintf("重启Deployment失败, %v\n", err))
	}
	return nil
}

// CreateDeployment 创建 Deployment
func (d *deployment) CreateDeployment(client *kubernetes.Clientset, data *vo.DeploymentCreate) (err error) {
	deploymentData := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:         data.Name,
			GenerateName: "",
			Namespace:    data.Namespace,
			Labels:       data.Label,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &data.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: data.Label,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   data.Name,
					Labels: data.Label,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  data.Name,
							Image: data.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
		Status: appsv1.DeploymentStatus{},
	}

	//判断是否打开健康检查功能，若打开，则定义ReadinessProbe和LivenessProbe
	if data.HealthCheck {
		deploymentData.Spec.Template.Spec.Containers[0].ReadinessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthPath,
					//intstr.IntOrString的作用是端口可以定义为整行，也可以定义为字符串
					//type=0表示整行，使用intVal
					//type=1表示字符串，使用strVal
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			//初始化等待时间
			InitialDelaySeconds: 5,
			//超时时间
			TimeoutSeconds: 15,
			//执行间隔
			PeriodSeconds: 5,
		}

		deploymentData.Spec.Template.Spec.Containers[0].LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: data.HealthPath,
					//intstr.IntOrString的作用是端口可以定义为整行，也可以定义为字符串
					//type=0表示整行，使用intVal
					//type=1表示字符串，使用strVal
					Port: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			//初始化等待时间
			InitialDelaySeconds: 5,
			//超时时间
			TimeoutSeconds: 15,
			//执行间隔
			PeriodSeconds: 5,
		}
	}

	//定义容器的limit和request资源
	deploymentData.Spec.Template.Spec.Containers[0].Resources.Limits =
		map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    resource.MustParse(data.Cpu),
			corev1.ResourceMemory: resource.MustParse(data.Memory),
		}

	deploymentData.Spec.Template.Spec.Containers[0].Resources.Requests =
		map[corev1.ResourceName]resource.Quantity{
			corev1.ResourceCPU:    resource.MustParse(data.Cpu),
			corev1.ResourceMemory: resource.MustParse(data.Memory),
		}

	//创建deployment
	_, err = client.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deploymentData, metav1.CreateOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("创建 Deployment 失败, %v\n", err))
		return errors.New(fmt.Sprintf("创建 Deployment 失败, %v\n", err))
	}
	return nil
}
