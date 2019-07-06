package cmd

import (
	"log"

	"github.com/eyalkenig/meizam-api/api/app/migrate"
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql"

	"github.com/spf13/cobra"
)

func NewMigrateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "migrate",
		Short: "migrate runs MySQL migrations",
		Run:   getMigrateRunFunction(),
	}

	return command
}

func getMigrateRunFunction() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		mysqlBuilder := mysql.NewBuilder()
		mysql, err := mysqlBuilder.Build()

		if err != nil {
			log.Println("Failed creating a new MySQL instance: " + err.Error())
			return
		}

		err = migrate.Up(mysql)
		if err != nil {
			log.Println("Failed running migrate command: " + err.Error())
			panic(err.Error())
		}
	}
}
