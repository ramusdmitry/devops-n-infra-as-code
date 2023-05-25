package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	groupApp "group-app-service"
)

const (
	usersTable = "users"
)

func NewPostgresDB(cfg groupApp.DBConfig) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode))

	if err != nil {
		logrus.Fatalf("[DB] failed to connect to db\n cause: %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
