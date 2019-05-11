package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/eyalkenig/meizam-api/api/app/repository/mysql/models"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
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

func (mysql *MySQL) ListTeams(limit, offset int) ([]*models.Team, error) {
	teams, err := models.Teams(
		Limit(limit),
		Offset(offset),
	).All(context.Background(), mysql.db)

	if err != nil {
		log.Println(fmt.Sprintf("Failed listing teams: %s, limit: %d, offset: %d", err.Error(), limit, offset))
		return nil, err
	}
	return teams, nil
}
