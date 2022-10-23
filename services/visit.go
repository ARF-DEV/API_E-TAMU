package services

import (
	"E-TamuAPI/helpers"
	"E-TamuAPI/models"
	"E-TamuAPI/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type OTPTokenBody struct {
	Token string `json:"otp_token"`
}
type DateRange struct {
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

func GetVisitByID(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v models.Visit

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&v)

		if err != nil {
			fmt.Println("Error while getting visit by id: ", err.Error())
			helpers.ErrorResponseJSON(w, "Bad Request", http.StatusBadRequest)
			return
		}

		visit, err := visitRepo.GetVisitByID(v.VisitId)

		if err != nil {
			fmt.Println("Error while getting visit by id: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		helpers.SuccessResponseJSON(w, "Success", visit)
	}
}

func GetVisitByStaffID(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var u models.User

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			log.Println("Error on visits by Staff ID: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		visits, err := visitRepo.GetVisitByStaffID(u.UserId)

		if err != nil {
			fmt.Println("Error while getting visits by staff id: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", visits)
	}
}

func GetVisitsByDateRange(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d DateRange

		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(&d)

		if err != nil {
			log.Println("Error on visits by date range: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		visits, err := visitRepo.GetVisitByDate(d.StartDate, d.EndDate)

		if err != nil {
			fmt.Println("Error while getting visits by date range: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", visits)
	}
}
func Register(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v models.Visit

		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			log.Println("Error on Register visit: ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		validate := validator.New()

		err = validate.Struct(v)

		if err != nil {
			log.Println("Json Is Invalid", err.Error())
			helpers.ErrorResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		expTime := time.Now().Add(time.Minute * 10)
		tokenString, key, err := helpers.GenerateRegisterOTPClaims(v, expTime)

		if err != nil {
			log.Println("Error on Register visit: ", err.Error())
			helpers.ErrorResponseJSON(w, "Failed Generating Token", http.StatusInternalServerError)
			return
		}

		//Kirim kode OTP ke email
		err = helpers.SendOTPEmail(v.GuestEmail, key)

		if err != nil {
			log.Println("Error while sending Email", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Generating OTP Token", OTPTokenBody{
			Token: tokenString,
		})
	})
}
func VerifyOTPRegisterVisit(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var t OTPTokenBody

		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			log.Println("Error while parsing JSON: ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if len(t.Token) < 1 {
			log.Println("Error : OTP Token Guess is Empty", http.StatusBadRequest)
			helpers.ErrorResponseJSON(w, "OTP Token Guess Must not be Empty", http.StatusBadRequest)
			return
		}
		AuthHeader := r.Header.Get("Authorization")

		if !strings.Contains(AuthHeader, "OTP") {
			log.Println("Error while getting OTP token : OTP Token Not Found")
			helpers.ErrorResponseJSON(w, "OTP Token Not Found", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(AuthHeader, "OTP ", "", -1)

		claims := helpers.RegisterVisitOTPClaims{}

		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return []byte(os.Getenv("JWT_OTP_KEY")), nil
		})

		if err != nil {
			log.Println("Error while parsing claims: ", err.Error())
			helpers.ErrorResponseJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("Error : Token is Invalid")
			helpers.ErrorResponseJSON(w, "Unautherized", http.StatusUnauthorized)
			return
		}

		if t.Token != claims.OTPSecret {
			log.Println("Error : OTP Token didn't Match")
			helpers.ErrorResponseJSON(w, "Invalid OTP Token", http.StatusUnauthorized)
			return
		}

		_, err = visitRepo.CreateVisit(claims.VisitData)

		if err != nil {
			log.Println("Error while inserting user: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Registering Visit Using OTP", claims.VisitData)

	})
}
func GetAllVisit(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		visits, err := visitRepo.GetAllVisit()

		if err != nil {
			fmt.Println("Error while Getting all Visit: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", visits)
	}
}

func DeleteVisitByID(visitRepo *repository.VisitRepository) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			fmt.Println("Error while deleting Visit: URL param is not found")
			helpers.ErrorResponseJSON(w, "Bad Request", http.StatusBadRequest)
			return
		}

		idInt, err := strconv.Atoi(id)

		if err != nil {
			fmt.Println("Error while parsing id param: ", err.Error())
			helpers.ErrorResponseJSON(w, "id is invalid", http.StatusBadRequest)
			return
		}

		Visit, err := visitRepo.DeleteVisit(idInt)

		if err != nil {
			fmt.Println("Error while deleting Visit: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", Visit)
	}
}
