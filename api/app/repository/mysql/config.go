package mysql

type Config struct {
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port     string `envconfig:"MYSQL_PORT" default:"3306"`
	Username string `envconfig:"MYSQL_USERNAME" default:"root"`
	Password string `envconfig:"MYSQL_PASSWORD" default:"password"`
	Database string `envconfig:"MYSQL_DATABASE" default:"meizamDB"`
}
