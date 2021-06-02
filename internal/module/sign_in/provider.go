package sign_in

import (
	"github.com/google/wire"
	"gomod/internal/module/sign_in/model"
)

var Provider = wire.NewSet(
	model.ModelProvider,
	ServiceProvider,
)

