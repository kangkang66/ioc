package model

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gomod/internal/module/task/model/key"
)

// 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象
func ModelProvider(rdsHandle *redis.Client,dbHandle	*gorm.DB) (*Model,error) {
	return &Model{
		rdsHandle: rdsHandle,
		dbHandle: dbHandle,
	},nil
}

type Model struct {
	rdsHandle	*redis.Client
	dbHandle	*gorm.DB
}

func (t *Model) GetTask(ctx context.Context, taskId string) (num int64,err error) {
	num,err = t.rdsHandle.Get(key.RdsTaskNum(taskId)).Int64()
	return
}