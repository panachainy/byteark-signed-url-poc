//go:build wireinject
// +build wireinject

//go:generate wire
package app

import (
	"go-boilerplate/cmd/config"
	"go-boilerplate/cmd/log"

	"github.com/google/wire"
)

func Wire() (*Application, error) {
	wire.Build(
		Provide,
		config.ProviderSet,
		log.ProviderSet,
	)

	return &Application{}, nil
}
