package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	postApp "posts-app-service/pkg/model"
	"strings"
	"time"
)

type PostListPostgres struct {
	db *sqlx.DB
}

func NewPostListPostgres(db *sqlx.DB) *PostListPostgres {
	return &PostListPostgres{db: db}
}

func (r *PostListPostgres) CreatePost(userId int, post postApp.Post) (int, error) {

	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var postId int

	createPostQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, description) VALUES ($1, $2, $3) RETURNING id", postsTable)
	row := tx.QueryRow(createPostQuery, userId, post.Title, post.Description)

	if err := row.Scan(&postId); err != nil {
		logrus.Errorf("[%s] [DB] failed to insert new comment for post (%d) into db",
			time.Now().UTC().Format("2006-01-02 15:04:05"), postId)
		tx.Rollback()
		return 0, err
	}

	logrus.Infof("[%s] [DB] post (%d) by user (%d) successfully inserted into db",
		time.Now().UTC().Format("2006-01-02 15:04:05"), postId, userId)

	return postId, tx.Commit()

}

func (r *PostListPostgres) GetAllPosts() ([]postApp.Post, error) {
	var posts []postApp.Post

	query := fmt.Sprintf("SELECT p.id, u.username, p.user_id, p.title, p.description FROM %s p LEFT JOIN users u ON p.user_id=u.id", postsTable)
	err := r.db.Select(&posts, query)
	return posts, err
}

func (r *PostListPostgres) GetAllComments(postId int) ([]postApp.Comment, error) {
	var comments []postApp.Comment

	query := fmt.Sprintf("SELECT id, post_id, user_id, user_name, content FROM %s WHERE post_id=$1", commsTable)
	err := r.db.Select(&comments, query, postId)
	return comments, err
}

func (r *PostListPostgres) UpdatePostById(userId, postId int, input postApp.UpdatePostInput) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d AND user_id=$%d", postsTable, setQuery, argId, argId+1)
	args = append(args, postId, userId)

	logrus.Warnf("updateQuery: %s", query)
	logrus.Warnf("args: %s", args)

	_, err = tx.Exec(query, args...)

	if err != nil {
		logrus.Warnf("[%s] [DB] failed to update post (%d) by user (%d), cause: %s",
			time.Now().UTC().Format("2006-01-02 15:04:05"),
			postId, userId, err.Error())
		tx.Rollback()
		return err
	}

	return err

}

func (r *PostListPostgres) DeletePostById(userId, postId int) error {
	query := fmt.Sprintf("DELETE FROM %s p USING %s u WHERE (p.user_id=$1 AND p.id=$2) OR ((u.group_id=1 OR u.group_id=2) AND p.id=$2)", postsTable, usersTable)
	fmt.Println(query)
	result, err := r.db.Exec(query, userId, postId)

	deletedRows, _ := result.RowsAffected()

	if deletedRows == 0 {
		return errors.New(fmt.Sprintf("cannot delete postId=%d for userID=%d", postId, userId))
	}

	logrus.Infof("[%s] [DB] post (%d) by user (%d) successfully deleted",
		time.Now().UTC().Format("2006-01-02 15:04:05"), postId, userId)

	return err
}
