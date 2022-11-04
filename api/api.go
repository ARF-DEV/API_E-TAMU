package api

import (
	"E-TamuAPI/middleware"
	"E-TamuAPI/repository"
	"E-TamuAPI/services"
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

	r.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))).ServeHTTP)

	r.Group(func(r chi.Router) {
		r.Get("/api/v1/login", services.UserLogin(a.UserRepo))
		r.Post("/api/v1/visit/create", services.RegisterVisit(a.VisitRepo))
		r.Post("/api/v1/visit/verify", services.VerifyOTPRegisterVisit(a.VisitRepo, a.UserRepo))
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authorization(a.UserRepo))
		r.Get("/api/v1/visits", services.GetAllVisit(a.VisitRepo))
		r.Get("/api/v1/visits/{id}", services.GetVisitByID(a.VisitRepo))
		r.Get("/api/v1/visits/date", services.GetVisitsByDateRange(a.VisitRepo))
		r.Get("/api/v1/visits/staff/{id}", services.GetVisitByStaffID(a.VisitRepo))
		r.Delete("/api/v1/visits/{id}", services.DeleteVisitByID(a.VisitRepo))
		r.Get("/api/v1/users", services.GetAllUser(a.UserRepo))
		r.Get("/api/v1/user/{id}", services.GetUserByID(a.UserRepo))
		r.Put("/api/v1/user", services.UpdateUserByID(a.UserRepo))
		r.Post("/api/v1/user", services.CreateUser(a.UserRepo))
		r.Delete("/api/v1/users/{id}", services.DeleteUserByID(a.UserRepo))
		r.Get("/api/v1/visits/generate", services.GenerateFileVisit(a.VisitRepo, a.UserRepo))
		r.Post("/api/v1/visits/{id}/confirm", services.ConfirmVisit(a.VisitRepo))
		// Export Data

	})
	return r
}
