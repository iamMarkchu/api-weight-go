package main

import (
	"api-weight-go/internal/conf"
	"api-weight-go/pkg/auth"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	id := uint64(668)
	token := getTestAuthKey(id)
	fmt.Println(token)
}

func getTestAuthKey(id uint64) string {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	key := bc.Auth.GetKey()
	claims := &auth.ApiClaims{
		id,
		jwtv4.RegisteredClaims{
			NotBefore: jwtv4.NewNumericDate(time.Now()),
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(time.Second * 86400)),
			Issuer:    key,
		},
	}
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(key))
	return t
}
