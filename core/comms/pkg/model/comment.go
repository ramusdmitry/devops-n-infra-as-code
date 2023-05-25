package model

import "errors"

type Comment struct {
	Id       int    `json:"id" db:"id"`
	UserId   int    `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	PostId   int    `json:"post_id" db:"post_id" binding:"required"`
	Content  string `json:"content" db:"content" binding:"required"`
}

type UpdateCommentInput struct {
	Content *string `json:"content" binding:"required"`
}

func (i UpdateCommentInput) Validate() error {
	if i.Content == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
