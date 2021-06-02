package task

import (
	"context"
	"gomod/internal/module/task/model"
)

func ServiceProvider(modelHandle *model.Model) (*Service,error) {
	return &Service{modelHandle: modelHandle},nil
}

type Service struct {
	modelHandle *model.Model
}
func (t *Service) Info(ctx context.Context, taskId string) (data map[string]interface{},err error) {
	num,err := t.modelHandle.GetTask(ctx, taskId)
	if err != nil {
		return
	}
	data = map[string]interface{}{
		"num":num,
	}
	return
}
