// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/iamMarkchu/api-weight-go/internal/biz"
	"github.com/iamMarkchu/api-weight-go/internal/conf"
	"github.com/iamMarkchu/api-weight-go/internal/data"
	"github.com/iamMarkchu/api-weight-go/internal/server"
	"github.com/iamMarkchu/api-weight-go/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Auth, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
