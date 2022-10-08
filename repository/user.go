package repository

import (
	"E-TamuAPI/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAllUser() ([]models.User, error) {
	sqlStatement := `SELECT user_id, user_name, user_email, user_role, user_password FROM user_data;`

	rows, err := u.db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserRole, &user.UserPassword)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) GetUserByID(userId int) (*models.User, error) {
	sqlStatement := `
	SELECT user_id, user_name, user_email, user_role, user_password 
	FROM user_data 
	WHERE user_id = ?;`

	var user models.User
	err := u.db.QueryRow(sqlStatement, userId).Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserRole, &user.UserPassword)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	sqlStatement := `
	SELECT user_id, user_name, user_email, user_role, user_password 
	FROM user_data 
	WHERE user_email = ?;`

	var user models.User
	err := u.db.QueryRow(sqlStatement, email).Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserRole, &user.UserPassword)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(user models.User) (*models.User, error) {
	sqlStatement := `
	UPDATE user_data
	SET 
		user_name = ?,
		user_email = ?,
		user_role = ?,
		user_password = ?
	WHERE user_id = ? 
	RETURNING user_id, user_name, user_email, user_role, user_password;
	`
	var result models.User
	err := u.db.QueryRow(sqlStatement, user.UserName, user.UserEmail, user.UserRole, user.UserPassword, user.UserId).
		Scan(&result.UserId, &result.UserName, &result.UserEmail, &result.UserRole, &result.UserPassword)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *UserRepository) CreateUser(user models.User) (*models.User, error) {
	sqlStatement := `
	INSERT INTO user_data (user_name, user_email, user_role, user_password) 
	VALUES (?, ?, ?, ?)
	RETURNING user_id, user_name, user_email, user_role, user_password;
	`

	var result models.User
	err := u.db.QueryRow(sqlStatement, user.UserName, user.UserEmail, user.UserRole, user.UserPassword, user.UserId).
		Scan(&result.UserId, &result.UserName, &result.UserEmail, &result.UserRole, &result.UserPassword)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *UserRepository) DeleteUser(userId int) (*models.User, error) {
	sqlStatement := `
	DELETE FROM user_data 
	WHERE user_id = ?
	RETURNING user_id, user_name, user_email, user_role, user_password;
	`

	var result models.User
	err := u.db.QueryRow(sqlStatement, userId).
		Scan(&result.UserId, &result.UserName, &result.UserEmail, &result.UserRole, &result.UserPassword)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
