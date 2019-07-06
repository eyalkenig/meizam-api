package migrate

import (
	"fmt"
	"log"

	"github.com/eyalkenig/meizam-api/api/app/repository/mysql"
)

func Up(mysql *mysql.MySQL) error {
	log.Println("Running all pending migrations")

	appliedMigrations, err := mysql.Migrate()
	if err != nil {
		log.Println(fmt.Sprintf("Failed to apply migrations: %s", err.Error()))
		return err
	}
	log.Println(fmt.Sprintf("Finished running migrations successfully: %v", appliedMigrations))
	return nil
}
