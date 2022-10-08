package api

import (
	"E-TamuAPI/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type API struct {
	UserRepo  *repository.UserRepository
	VisitRepo *repository.VisitRepository
}

func NewAPI(db *sql.DB) *API {
	return &API{
		UserRepo:  repository.NewUserRepository(db),
		VisitRepo: repository.NewVisitRepository(db),
	}
}

func (a *API) GetRouter() http.Handler {
	r := chi.NewRouter()

	return r
}
