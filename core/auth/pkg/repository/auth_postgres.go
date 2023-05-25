package repository

import (
	authApp "auth-app-service/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	admin = iota + 1
	moderator
	journalist
	guest
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user authApp.User) (int, error) {
	var userId int
	query := fmt.Sprintf("INSERT INTO %s (name, group_id, username, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, journalist, user.Username, user.Password)
	if err := row.Scan(&userId); err != nil {
		logrus.Errorf("[%s] [DB]: (create user) failed to insert into %s\n cause: %s",
			time.Now().UTC().Format("2006-01-02 15:04:05"), usersTable, err.Error())
		return 0, err
	}
	logrus.Infof("[%s] [DB]: (create user) successfully inserted (%s, %s) into db with id=%d",
		time.Now().UTC().Format("2006-01-02 15:04:05"), user.Name, user.Username, userId)
	return userId, nil
}

func (r *AuthPostgres) GetUser(username, password string) (authApp.Profile, error) {
	var user authApp.Profile
	query := fmt.Sprintf("SELECT id, name, username, group_id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		logrus.Errorf("[%s] [DB] failed to get user from db with username=%s\n cause: %s",
			time.Now().UTC().Format("2006-01-02 15:04:05"), username, err.Error())
		return authApp.Profile{}, err
	}
	logrus.Infof("[%s] [DB] successfully get user (%d) from db with username=%s and ",
		time.Now().UTC().Format("2006-01-02 15:04:05"), user.Id, user.Username)
	return user, err
}

func (r *AuthPostgres) GetUserById(id int) (authApp.Profile, error) {
	var profile authApp.Profile
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&profile, query, id)
	if err != nil {
		logrus.Errorf("[%s] [DB] failed to get user by id (%d)\n cause: %s",
			time.Now().UTC().Format("2006-01-02 15:04:05"), id, err.Error())
		return authApp.Profile{}, err
	}
	logrus.Infof("[%s] [DB] successfully get user profile by id (%d)", time.Now().UTC().Format("2006-01-02 15:04:05"), id)
	return profile, err
}
