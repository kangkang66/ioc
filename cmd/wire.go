// +build wireinject

package main

import (
	"github.com/google/wire"
	"gomod/internal/api"
	"gomod/internal/module"
	"gomod/pkg"
)

func InitApi() (*api.Api,error) {
	panic(wire.Build(pkg.Provider, module.Provider, api.Provider))
}
