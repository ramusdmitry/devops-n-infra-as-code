package posts_app_service

import "errors"

//type Post struct {
//	PostId   int       `json:"postId"`
//	UserId   int       `json:"userId"`
//	Username string    `json:"username"`
//	Content  string    `json:"content"`
//	Time     time.Time `json:"time"`
//}

type Post struct {
	Id          int      `json:"id" db:"id"` //post id
	UserId      int      `json:"userId" db:"user_id"`
	UserName    string   `json:"username" db:"username"`
	Title       string   `json:"title" db:"title" binding:"required"`
	Description string   `json:"description" db:"description"`
	Comments    Comments `json:"comms"`
}

type Comments struct {
	Data []Comment `json:"data"`
}

type Comment struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	PostId   int    `json:"post_id"`
	Content  string `json:"content"`
}

type UpdatePostInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
