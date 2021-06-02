package model

import (
	"context"
	"github.com/go-redis/redis"
	"gomod/internal/module/task/model/key"
)

// 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象
func ModelProvider(rdsHandle *redis.Client) (*Model,error) {
	return &Model{
		rdsHandle: rdsHandle,
	},nil
}

type Model struct {
	rdsHandle	*redis.Client
}

func (t *Model) GetTask(ctx context.Context, taskId string) (num int64,err error) {
	num,err = t.rdsHandle.Get(key.RdsTaskNum(taskId)).Int64()
	return
}