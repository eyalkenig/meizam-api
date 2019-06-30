package auth

type Config struct {
	Audience string `envconfig:"AUTH0_AUDIENCE" default:""`
	Domain   string `envconfig:"AUTH0_DOMAIN" default:"localhost"`
}
