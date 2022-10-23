package models

type User struct {
	UserId       int    `db:"user_id" json:"user_id"`
	UserName     string `db:"user_name" json:"user_name"`
	UserEmail    string `db:"user_email" json:"user_email"`
	UserRole     string `db:"user_role" json:"user_role"`
	UserPassword string `db:"user_password" json:"user_password"`
}
