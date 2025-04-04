package dao

import (
	"KubeEase/db"
	"KubeEase/model/po"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

var Event event

type event struct{}

type Events struct {
	Items []*po.Event `json:"items"`
	Total int         `json:"total"`
}

// GetList 获取列表
func (*event) GetList(name, cluster string, page, size int) (*Events, error) {
	//定义分页数据的起始位置
	startSet := (page - 1) * size
	//定义数据库查询的返回内容
	var (
		eventList = make([]*po.Event, 0)
		total     = 0
	)

	//数据库查询，先查total
	tx := db.GORM.Model(&po.Event{}).
		Where("name like ? and cluster = ?", "%"+name+"%", cluster).
		Count(&total)

	if tx.Error != nil {
		zap.L().Error("获取Event列表失败," + tx.Error.Error())
		return nil, errors.New("获取Event列表失败," + tx.Error.Error())
	}

	//数据库查询
	tx = db.GORM.Model(&po.Event{}).
		Where("name like ? and cluster = ?", "%"+name+"%", cluster).
		Limit(size).
		Offset(startSet).
		Order("id desc").
		Find(&eventList)

	if tx.Error != nil {
		zap.L().Error(fmt.Sprintf("获取Event列表失败, %v\n", tx.Error))
		return nil, errors.New(fmt.Sprintf("获取Event列表失败, %v\n", tx.Error))
	}

	return &Events{
		Items: eventList,
		Total: total,
	}, nil
}

// Add 新增
func (*event) Add(event *po.Event) error {
	tx := db.GORM.Create(&event)
	if tx.Error != nil {
		zap.L().Error(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
		return errors.New(fmt.Sprintf("添加Event失败, %v\n", tx.Error))
	}
	return nil
}

// HasEvent 查询单个 event
func (*event) HasEvent(name, kind, namespace, reason string, eventTime time.Time, cluster string) (*po.Event, bool, error) {
	data := &po.Event{}
	//tx := db.GORM.Where("name = ? and kind = ? and namespace = ? and reason = ? and event_time = ? and cluster = ?",
	//	name, kind, namespace, reason, eventTime, cluster).First(&data)
	tx := db.GORM.Where("name = ? and kind = ? and namespace = ? and reason = ? and event_time = ? and cluster = ?",
		name, kind, namespace, reason, eventTime, cluster).First(&data)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, nil
	}
	if tx.Error != nil {
		zap.L().Error(fmt.Sprintf("查询Event失败, %v\n", tx.Error))
		return nil, false, errors.New(fmt.Sprintf("查询Event失败, %v\n", tx.Error))
	}
	return data, true, nil
}
