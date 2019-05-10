package mysql

import (
	"fmt"
	"github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"
	"github.com/volatiletech/null"
	"log"

	"github.com/volatiletech/sqlboiler/boil"

	"context"
	"database/sql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error) {
	team := &models.Team{Name: teamName, ExternalEntityID: null.StringFromPtr(externalEntityId), ImageURL: null.StringFromPtr(imageUrl)}
	err := team.Insert(context.Background(), mysql.db, boil.Infer())
	if err != nil {
		log.Println(fmt.Sprintf("Failed creating unsubscriber: %s", err.Error()))
		return nil, err
	}
	return team, nil
}
