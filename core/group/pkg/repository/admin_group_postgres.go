package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"group-app-service/pkg/model"
	"strings"
)

type AdminGroupPostgres struct {
	db *sqlx.DB
}

func (r *AdminGroupPostgres) GetAllUsers() ([]model.User, error) {
	var users []model.User

	query := fmt.Sprintf("SELECT id, name, username, group_id FROM %s ORDER BY id ASC", usersTable)
	err := r.db.Select(&users, query)
	return users, err
}

func (r *AdminGroupPostgres) UpdateUsers(users model.UpdateUsers) error {

	queries := make([]string, 0)

	for _, user := range users.Data {
		q := fmt.Sprintf("UPDATE users SET group_id=%d WHERE id=%d;", user.GroupId, user.Id)
		queries = append(queries, q)

	}

	query := strings.Join(queries, "")

	_, err := r.db.Exec(query)

	if err != nil {
		fmt.Errorf("%s", err.Error())
		return err
	}

	return err

}

func (r *AdminGroupPostgres) DeleteUsers() error {
	q := "DELETE FROM comments;"
	if _, err := r.db.Exec(q); err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	q = "DELETE FROM posts;"
	if _, err := r.db.Exec(q); err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	q = "DELETE FROM users;"
	if _, err := r.db.Exec(q); err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func NewAdminGroupPostgres(db *sqlx.DB) *AdminGroupPostgres {
	return &AdminGroupPostgres{db: db}
}
