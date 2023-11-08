//go:build wireinject
// +build wireinject

//go:generate wire
package config

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	Provide,
)

func Wire() (*Configuration, error) {
	wire.Build(ProviderSet)
	return &Configuration{}, nil
}
