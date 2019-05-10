package mysql

import (
	"database/sql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) CreateTeam(teamName, externalEntityId, imageUrl string) (int, error) {
	return -1, nil
}
