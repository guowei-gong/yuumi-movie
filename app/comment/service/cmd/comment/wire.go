// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"yuumi-movie/app/comment/service/internal/biz"
	"yuumi-movie/app/comment/service/internal/conf"
	"yuumi-movie/app/comment/service/internal/data"
	"yuumi-movie/app/comment/service/internal/server"
	"yuumi-movie/app/comment/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
