package model

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Profile struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"user_name" db:"username"`
	GroupID  int    `json:"group_id" db:"group_id"`
}
