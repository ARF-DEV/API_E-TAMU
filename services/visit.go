package services

import (
	"E-TamuAPI/helpers"
	"E-TamuAPI/models"
	"E-TamuAPI/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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
