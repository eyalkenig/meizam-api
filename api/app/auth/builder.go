package auth

import (
	"log"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/kelseyhightower/envconfig"
)

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (f *Builder) Build() (*jwtmiddleware.JWTMiddleware, error) {
	conf := &Config{}
	if err := loadConfigObject(conf); err != nil {
		log.Fatalf("failed loading configuration for auth: %s", err.Error())
		return nil, err
	}
	return GetJwtMiddleware(conf.Domain, conf.Audience), nil
}

func loadConfigObject(configObject interface{}) error {
	return envconfig.Process("", configObject)
}
