package model

type User struct {
	UserId   string `json:"user_id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
