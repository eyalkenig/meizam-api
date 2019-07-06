package mysql

import (
	"context"
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
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

func (mysql *MySQL) Migrate() (int, error) {
	migrate.SetTable("schema_migrations")
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	return migrate.Exec(mysql.db, "mysql", migrations, migrate.Up)
}

func (mysql *MySQL) CreateTeam(teamName string, externalEntityId, imageUrl *string) (*models.Team, error) {
	team := &models.Team{Name: teamName, ExternalEntityID: null.StringFromPtr(externalEntityId), ImageURL: null.StringFromPtr(imageUrl)}
	err := team.Insert(context.Background(), mysql.db, boil.Infer())
	if err != nil {
		log.Println(fmt.Sprintf("Failed creating team: %s", err.Error()))
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

func (mysql *MySQL) CreateCompetition(name, competitionType string, externalEntityId *string) (*models.Competition, error) {
	competition := &models.Competition{Name: name, Type: competitionType, ExternalEntityID: null.StringFromPtr(externalEntityId)}
	err := competition.Insert(context.Background(), mysql.db, boil.Infer())
	if err != nil {
		log.Println(fmt.Sprintf("Failed creating competition: %s", err.Error()))
		return nil, err
	}
	return competition, nil
}

func (mysql *MySQL) ListCompetitions(limit, offset int) ([]*models.Competition, error) {
	competitions, err := models.Competitions(
		Limit(limit),
		Offset(offset),
	).All(context.Background(), mysql.db)

	if err != nil {
		log.Println(fmt.Sprintf("Failed listing competitions: %s, limit: %d, offset: %d", err.Error(), limit, offset))
		return nil, err
	}
	return competitions, nil
}
