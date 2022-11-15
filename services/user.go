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

type userForVisit struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

func GetUserByToken(userRepo *repository.UserRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userData, ok := r.Context().Value("user_data").(models.User)
		if !ok {
			log.Println("userData not found")
			helpers.ErrorResponseJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Getting User By Token", userData)

	})
}
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
		id := chi.URLParam(r, "id")

		if len(id) < 1 {
			log.Println("Error product by id : id query not found")
			helpers.ErrorResponseJSON(w, "id query required", http.StatusBadRequest)
			return
		}

		id_int, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id query", http.StatusBadRequest)
			return
		}
		user, err := userRepo.GetUserByID(id_int)

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
		name := r.URL.Query().Get("name")

		if len(name) < 1 {
			log.Println("Error : name query not found")
			helpers.ErrorResponseJSON(w, "name query required", http.StatusBadRequest)
			return
		}

		users, err := userRepo.GetUsersByName(name)

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
		fmt.Println(u)

		user, err := userRepo.GetUserByEmail(u.UserEmail)

		if err != nil {
			log.Println("Error on Login: ", err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "User Not Found", http.StatusNotFound)
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
		name := r.URL.Query().Get("name")

		if len(name) > 0 {
			GetUsersByName(userRepo).ServeHTTP(w, r)
		} else {
			users, err := userRepo.GetAllUser()

			if err != nil {
				fmt.Println("Error while getting all user : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			helpers.SuccessResponseJSON(w, "Success", users)
		}

	}
}

func GetVisitedUser(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if len(id) < 1 {
			log.Println("Error product by id : id query not found")
			helpers.ErrorResponseJSON(w, "id query required", http.StatusBadRequest)
			return
		}

		id_int, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id query", http.StatusBadRequest)
			return
		}
		user, err := userRepo.GetUserByID(id_int)

		if err != nil {
			fmt.Println("Error while getting users by name : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", userForVisit{UserName: user.UserName, UserID: user.UserId})
	}
}

func GetAllAvailableUser(userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")

		if len(name) > 0 {
			fmt.Println(name)
			users, err := userRepo.GetUsersByName(name)

			if err != nil {
				fmt.Println("Error while getting users by name : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var response []userForVisit
			for _, user := range users {
				var u userForVisit

				u.UserID = user.UserId
				u.UserName = user.UserName

				response = append(response, u)
			}

			helpers.SuccessResponseJSON(w, "Success", response)
		} else {
			users, err := userRepo.GetAllUser()

			if err != nil {
				fmt.Println("Error while getting all user : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			var response []userForVisit
			for _, user := range users {
				var u userForVisit

				u.UserID = user.UserId
				u.UserName = user.UserEmail

				response = append(response, u)
			}

			helpers.SuccessResponseJSON(w, "Success", response)

		}

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
