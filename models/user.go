package models

type User struct {
	UserId       int    `db:"user_id" json:"user_id" validate:"required"`
	UserName     string `db:"user_name" json:"user_name" validate:"required"`
	UserEmail    string `db:"user_email" json:"user_email" validate:"required,email"`
	UserRole     string `db:"user_role" json:"user_role" validate:"required"`
	UserPassword string `db:"user_password" json:"user_password" validate:"required"`
}
