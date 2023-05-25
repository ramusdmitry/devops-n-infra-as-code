package model

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	GroupId  int    `json:"group_id" db:"group_id"`
}

type UpdateUsers struct {
	Data []User `json:"data"`
}
