// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"api-weight-go/internal/biz"
	"api-weight-go/internal/conf"
	"api-weight-go/internal/data"
	"api-weight-go/internal/server"
	"api-weight-go/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Auth, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
