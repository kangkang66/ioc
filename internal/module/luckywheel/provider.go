package luckywheel

import (
	"github.com/google/wire"
	"gomod/internal/module/luckywheel/model"
)

var Provider = wire.NewSet(
	model.ModelProvider,
	ServiceProvider,
)

