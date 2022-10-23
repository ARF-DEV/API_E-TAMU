package services

import (
	"E-TamuAPI/helpers"
	"E-TamuAPI/models"
	"E-TamuAPI/repository"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func UpdateUserByID(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on Create: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		validate := validator.New()

		err = validate.Struct(u)

		if err != nil {
			log.Println("user is not valid : ", err.Error())
			helpers.ErrorResponseJSON(w, "user invalid : "+err.Error(), http.StatusBadRequest)
			return
		}

		updatedUser, err := userRepo.UpdateUser(u)

		if err != nil {
			log.Println("Error on update user repo : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", updatedUser)
	}
}
func CreateUser(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on Create: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		user, err := userRepo.CreateUser(u)

		if err != nil {
			log.Println("Error on Create: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", user)
	}
}

func GetUserByID(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on get users by name: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		user, err := userRepo.GetUserByID(u.UserId)

		if err != nil {
			fmt.Println("Error while getting users by name : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", user)
	}
}
func GetUsersByName(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on get users by name: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		users, err := userRepo.GetUsersByName(u.UserName)

		if err != nil {
			fmt.Println("Error while getting users by name : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", users)
	}
}

func UserLogin(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on Login: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		user, err := userRepo.GetUserByEmail(u.UserEmail)

		if err != nil {
			log.Println("Error on Login: ", err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "User Not Found", http.StatusOK)
				return
			}

			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if user.UserPassword != u.UserPassword {
			log.Println("Error on Login: Password is invalid")
			helpers.ErrorResponseJSON(w, "Password is invalid", http.StatusUnauthorized)
			return
		}

		expTime := time.Now().Add(time.Minute * 30)
		tokenString, err := helpers.GenerateUserToken(*user, expTime)

		if err != nil {
			log.Println("Error on Login: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Login Success", map[string]interface{}{
			"token": tokenString,
		})

	}
}

func GetAllUser(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userRepo.GetAllUser()

		if err != nil {
			fmt.Println("Error while getting all user : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", users)
	}
}

func DeleteUserByID(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			fmt.Println("Error while deleting User: URL param is not found")
			helpers.ErrorResponseJSON(w, "Bad Request", http.StatusBadRequest)
			return
		}

		idInt, err := strconv.Atoi(id)

		if err != nil {
			fmt.Println("Error while parsing id param: ", err.Error())
			helpers.ErrorResponseJSON(w, "id is invalid", http.StatusBadRequest)
			return
		}

		user, err := userRepo.DeleteUser(idInt)

		if err != nil {
			fmt.Println("Error while deleting user: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", user)
	}
}
