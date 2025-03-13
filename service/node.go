package service

import (
	"KubeEase/model/vo"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var Node node

type node struct{}

// toCells 方法用于将Node类型数组，转换成DataCell类型数组
func (n *node) toCells(std []corev1.Node) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = nodeCell(std[i])
	}
	return cells
}

// fromCells 方法用于将DataCell类型数组，转换成Node类型数组
func (n *node) fromCells(cells []DataCell) []corev1.Node {
	nodes := make([]corev1.Node, len(cells))
	for i := range cells {
		nodes[i] = corev1.Node(cells[i].(nodeCell))
	}
	return nodes
}

// GetNodes 获取 Node 列表
func (n *node) GetNodes(client *kubernetes.Clientset, filterName string, size, page int) (nodesResp *vo.NodesResp, err error) {
	nodeList, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("获取 Node 列表失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取 Node 列表失败, %v\n", err))
	}
	selectableData := &dataSelector{
		GenericDataList: n.toCells(nodeList.Items),
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
	//将[]DataCell类型的Service列表转为v1.Service列表
	nodes := n.fromCells(data.GenericDataList)
	return &vo.NodesResp{
		Items: nodes,
		Total: total,
	}, nil
}

// GetNodeDetail 获取 Node 详情
func (n *node) GetNodeDetail(client *kubernetes.Clientset, nodeName string) (node *corev1.Node, err error) {
	node, err = client.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		zap.L().Error(fmt.Sprintf("获取 Node 详情失败, %v\n", err))
		return nil, errors.New(fmt.Sprintf("获取 Node 详情失败, %v\n", err))
	}
	return node, nil
}
