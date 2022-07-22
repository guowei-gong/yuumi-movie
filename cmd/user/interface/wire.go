//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"yuumi-movie/internal/biz"
	"yuumi-movie/internal/conf"
	"yuumi-movie/internal/data"
	"yuumi-movie/internal/server"
	"yuumi-movie/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.JWT, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
