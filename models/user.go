package models

import "E-TamuAPI/schema"

type User struct {
	UserId       int    `db:"user_id"`
	UserName     string `db:"user_name"`
	UserEmail    string `db:"user_email"`
	UserRole     string `db:"user_role"`
	UserPassword string `db:"user_password"`
}

func (a *User) GetSchema() schema.UserSchema {
	return schema.UserSchema{
		UserId:       a.UserId,
		UserName:     a.UserName,
		UserEmail:    a.UserEmail,
		UserRole:     a.UserRole,
		UserPassword: a.UserPassword,
	}
}
