package repository

import (
	commsApp "comms-app-service/pkg/model"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type CommsPostgres struct {
	db *sqlx.DB
}

func (r *CommsPostgres) DeleteCommentById(userId int, commentId int) error {

	//query := fmt.Sprintf("DELETE FROM %s WHERE (id=%d AND user_id =%d) OR (id =%d AND user_id IN (SELECT id FROM users WHERE group_id IN (1, 2)));", commsTable, commentId, userId, commentId)
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", commsTable)
	result, err := r.db.Exec(query, commentId)

	deletedRows, err := result.RowsAffected()
	if deletedRows == 0 {
		logrus.Warnf("[%s] [DB] failed to delete comment (%d) for user (%d)",
			time.Now().UTC().Format("2006-01-02 15:04:05"), commentId, userId)
		return errors.New(fmt.Sprintf("cannot delete comment (%d) for user (%d)", commentId, userId))
	}

	logrus.Infof("[%s] [DB] comment (%d) was successfully deleted from db",
		time.Now().UTC().Format("2006-01-02 15:04:05"), commentId)

	return err

}

func (r *CommsPostgres) CreateComment(userId int, comment commsApp.Comment) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		//logrus.Error("Не удалось начать транзакцию", err.Error())
		return 0, err
	}

	var userName string

	getUsernameQuery := fmt.Sprintf("SELECT username FROM %s WHERE id=%d", usersTable, userId)
	userRow := tx.QueryRow(getUsernameQuery)

	if err := userRow.Scan(&userName); err != nil {
		logrus.Errorf("[%s] [DB] failed to get username, cause: %s", time.Now().UTC().Format("2006-01-02 15:04:05"), err.Error())
		return 0, err
	}

	var commentId int
	createCommQuery := fmt.Sprintf("INSERT INTO %s (user_name, user_id, post_id, content) VALUES ($1, $2, $3, $4) RETURNING id", commsTable)

	row := tx.QueryRow(createCommQuery, userName, userId, comment.PostId, comment.Content)

	if err := row.Scan(&commentId); err != nil {
		logrus.Errorf("[%s] [DB] failed to insert new comment for post (%d) into db",
			time.Now().UTC().Format("2006-01-02 15:04:05"), comment.PostId)
		return 0, err
	}

	logrus.Infof("[%s] [DB] comment (%d) for post (%d) successfully inserted into db",
		time.Now().UTC().Format("2006-01-02 15:04:05"), commentId, comment.PostId)

	return commentId, tx.Commit()
}

func (r *CommsPostgres) GetCommentsByPostId(postId int) ([]commsApp.Comment, error) {

	var comments []commsApp.Comment

	query := fmt.Sprintf("SELECT id, post_id, user_id, user_name, content FROM %s WHERE post_id=$1", commsTable)
	err := r.db.Select(&comments, query, postId)
	return comments, err

}

func (r *CommsPostgres) GetAllComments() ([]commsApp.Comment, error) {

	var comments []commsApp.Comment

	query := fmt.Sprintf("SELECT id, post_id, user_id, user_name, content FROM %s", commsTable)
	err := r.db.Select(&comments, query)
	return comments, err

}

func (r *CommsPostgres) UpdateCommentById(userId int, commentId int, comment commsApp.UpdateCommentInput) error {
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET content=$1 WHERE id=%d AND user_id=%d", commsTable, commentId, userId)

	logrus.Warnf("updateQuery: %s", query)

	result, err := tx.Exec(query, comment.Content)

	updatedRows, _ := result.RowsAffected()

	if updatedRows == 0 {
		logrus.Warnf("[%s] [DB] failed to update comment (%d) by user (%d), cause: %s",
			time.Now().UTC().Format("2006-01-02 15:04:05"),
			commentId, userId, err.Error())
		tx.Rollback()
		return err
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err

}

func NewCommsPostgres(db *sqlx.DB) *CommsPostgres {
	return &CommsPostgres{db: db}
}
