package task

import (
	"github.com/google/wire"
	"gomod/internal/module/task/model"
)

var Provider = wire.NewSet(
	model.ModelProvider,
	ServiceProvider,
)

