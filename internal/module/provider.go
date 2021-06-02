package module

import (
	"github.com/google/wire"
	"gomod/internal/module/luckywheel"
	"gomod/internal/module/sign_in"
	"gomod/internal/module/task"
)

var Provider = wire.NewSet(
	task.Provider,
	sign_in.Provider,
	luckywheel.Provider,
)