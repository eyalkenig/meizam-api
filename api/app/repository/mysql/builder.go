package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (f *Builder) Build() (*MySQL, error) {
	conf := &Config{}
	if err := loadConfigObject(conf); err != nil {
		log.Fatalf("failed loading configuration for MySQL: %s", err.Error())
		return nil, err
	}

	connectionString := buildConnectionString(conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)

	log.Printf("MySQL configuration: '%s' database on %s:%s - %s", conf.Database, conf.Host, conf.Port, connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed opening the database: %s", err.Error())
	}
	log.Println("database opened successfully")

	return NewMySQL(db), nil
}

func buildConnectionString(username, password, host, port, database string) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, database)
}
func loadConfigObject(configObject interface{}) error {
	return envconfig.Process("", configObject)
}
