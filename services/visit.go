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
type VisitCSV struct {
	VisitId            int    `csv:"visit_id"`
	UserVisitedName    string `csv:"user_visited_name"`
	GuestName          string `csv:"guest_name"`
	GuestEmail         string `csv:"guest_email"`
	VisitIntention     string `csv:"visit_intention"`
	VaccineCertificate string `csv:"vaccine_certificate"`
	VisitStatus        string `csv:"visit_status"`
	GuestCount         int    `csv:"guest_count"`
	VisitDate          string `csv:"visit_date"`
	VisitHour          string `csv:"visit_hour"`
	Transportation     string `csv:"transportation"`
}

func GetVisitByID(visitRepo *repository.VisitRepository) http.HandlerFunc {
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
		visit, err := visitRepo.GetVisitByID(id_int)

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

		visits, err := visitRepo.GetVisitByStaffID(id_int)

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
func RegisterVisit(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var v models.Visit

		// err := json.NewDecoder(r.Body).Decode(&v)
		// if err != nil {
		// 	log.Println("Error on Register visit: ", err.Error())
		// 	helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
		// 	return
		// }

		v.GuestName = r.FormValue("guest_name")
		v.UserVisitedId, _ = strconv.Atoi(r.FormValue("user_visited_id"))
		v.GuestEmail = r.FormValue("guest_email")
		v.VisitIntention = r.FormValue("visit_intention")
		v.VisitStatus = r.FormValue("visit_status")
		v.GuestCount, _ = strconv.Atoi(r.FormValue("guest_count"))
		v.VisitDate = r.FormValue("visit_date")
		v.VisitHour = r.FormValue("visit_hour")
		v.Transportation = r.FormValue("transportation")
		imageFile, imageHeader, err := r.FormFile("vaccine_certificate")
		fileLoc := ""
		if err == nil {
			fileLoc = helpers.UploadImage(imageFile, imageHeader)
		}

		v.VaccineCertificate = fileLoc

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
func VerifyOTPRegisterVisit(visitRepo *repository.VisitRepository, userRepo *repository.UserRepository) http.HandlerFunc {
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
			helpers.ErrorResponseJSON(w, "Unauthorized", http.StatusUnauthorized)
			helpers.RemoveFile(r, claims.VisitData.VaccineCertificate)
			return
		}

		if t.Token != claims.OTPSecret {
			log.Println("Error : OTP Token didn't Match")
			helpers.ErrorResponseJSON(w, "Invalid OTP Token", http.StatusUnauthorized)
			helpers.RemoveFile(r, claims.VisitData.VaccineCertificate)
			return
		}

		createdVisit, err := visitRepo.CreateVisit(claims.VisitData)
		if err != nil {
			log.Println("Error while inserting user: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		visitedStaff, err := userRepo.GetUserByID(createdVisit.UserVisitedId)

		if err != nil {
			log.Println("Error while getting visited staff : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		go helpers.SendVisitID(createdVisit.GuestEmail, createdVisit.VisitId)

		go helpers.SendVisitNotif(visitedStaff.UserEmail, createdVisit.GuestName)

		helpers.SuccessResponseJSON(w, "Success Registering Visit Using OTP", createdVisit)

	})
}
func GetAllVisit(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")

		if len(status) > 0 {
			GetVisitByStatus(visitRepo).ServeHTTP(w, r)
		} else {
			visits, err := visitRepo.GetAllVisit()

			if err != nil {
				fmt.Println("Error while Getting all Visit: ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			helpers.SuccessResponseJSON(w, "Success", visits)
		}

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

func ConfirmVisit(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if id == "" {
			fmt.Println("Error while confirming arrival: URL param is not found")
			helpers.ErrorResponseJSON(w, "Bad Request", http.StatusBadRequest)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error while parsing id param: ", err.Error())
			helpers.ErrorResponseJSON(w, "id is invalid", http.StatusBadRequest)
			return
		}

		// 1 = konfirmasi kedatangan 2 = konfirmasi keluar
		status := r.URL.Query().Get("status")

		if status == "" {
			fmt.Println("Error while getting quert : Status Query Required")
			helpers.ErrorResponseJSON(w, "status query required", http.StatusBadRequest)
			return
		}
		var visit *models.Visit
		if status == "1" {
			visit, err = visitRepo.ConfirmArrival(idInt)
		} else if status == "2" {
			visit, err = visitRepo.ConfirmFinish(idInt)
		} else {
			fmt.Println("Error : invalid status query")
			helpers.ErrorResponseJSON(w, "status query invalid", http.StatusBadRequest)
			return

		}

		if err != nil {
			fmt.Println("Error while confirming : ", err.Error())
			helpers.ErrorResponseJSON(w, "internal server error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "success", visit)

	}
}
func GenerateFileVisit(visitRepo *repository.VisitRepository, userRepo *repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var d DateRange

		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(&d)

		if err != nil {
			log.Println("Error on visits by date range (csv): ", err.Error())
			helpers.ErrorResponseJSON(w, "Json is Invalid", http.StatusBadRequest)
			return
		}

		visits, err := visitRepo.GetVisitByDate(d.StartDate, d.EndDate)

		if err != nil {
			fmt.Println("Error while getting visits by date range (csv): ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var visitsCSV []VisitCSV

		for _, visit := range visits {
			var v VisitCSV
			v.VisitId = visit.VisitId
			v.GuestName = visit.GuestName
			v.GuestEmail = visit.GuestEmail
			v.VisitIntention = visit.VisitIntention
			v.VaccineCertificate = visit.VaccineCertificate
			v.VisitStatus = visit.VisitStatus
			v.GuestCount = visit.GuestCount
			v.VisitDate = visit.VisitDate
			v.VisitHour = visit.VisitHour
			v.Transportation = visit.Transportation

			user, err := userRepo.GetUserByID(visit.UserVisitedId)

			if err != nil {
				fmt.Println("Error while processing csv : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			v.UserVisitedName = user.UserName

			visitsCSV = append(visitsCSV, v)

		}

		fileLoc, err := helpers.SaveToCSV(visitsCSV)
		if err != nil {
			fmt.Println("Error while saving CSV: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", map[string]string{
			"csv_url": fileLoc,
		})

	}
}

func GetVisitByStatus(visitRepo *repository.VisitRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")

		if status == "" {
			fmt.Println("Error while getting quert : Status Query Required")
			helpers.ErrorResponseJSON(w, "status query required", http.StatusBadRequest)
			return
		}

		visits, err := visitRepo.GetVisitByStatus(status)

		if err != nil {
			fmt.Println("Error while getting visits by status: ", err.Error())
			helpers.ErrorResponseJSON(w, "internal server error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success", visits)
	}
}
