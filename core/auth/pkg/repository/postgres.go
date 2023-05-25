package repository

import (
	authApp "auth-app-service"
	_ "errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

const usersTable = "users"

func NewPostgresDB(cfg authApp.DBConfig) (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode))

	if err != nil {
		logrus.Fatalf("[%s] [DB] failed to connect to db\n cause: %s", time.Now().UTC().Format("2006-01-02 15:04:05"), err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
