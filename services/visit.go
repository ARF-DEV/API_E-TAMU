package services

import (
	"E-TamuAPI/helpers"
	"E-TamuAPI/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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
