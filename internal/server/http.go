package server

import (
	v1 "api-weight-go/api/helloworld/v1"
	api_weight "api-weight-go/api/weight"
	"api-weight-go/internal/conf"
	"api-weight-go/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ca *conf.Auth, greeter *service.GreeterService, weight *service.WeightService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ca.GetKey()), nil
			}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	api_weight.RegisterWeightHTTPServer(srv, weight)
	return srv
}
