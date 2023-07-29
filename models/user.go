package models

type User struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}
