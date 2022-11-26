package api

import (
	"E-TamuAPI/middleware"
	"E-TamuAPI/repository"
	"E-TamuAPI/services"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))).ServeHTTP)

	r.Group(func(r chi.Router) {

		r.Post("/api/v1/login", services.UserLogin(a.UserRepo))
		r.Post("/api/v1/visit/create", services.RegisterVisit(a.VisitRepo))
		r.Post("/api/v1/visit/verify", services.VerifyOTPRegisterVisit(a.VisitRepo, a.UserRepo))
		r.Get("/api/v1/visit/users", services.GetAllAvailableUser(a.UserRepo))
		r.Get("/api/v1/visit/users/{id}", services.GetVisitedUser(a.UserRepo))
		r.Get("/api/v1/visits/{id}", services.GetVisitByID(a.VisitRepo))
		r.Get("/api/v1/visits/confirmvisit", services.ConfirmVisitProposalRedirect(a.VisitRepo))
		r.Get("/api/v1/visits/cancelvisit", services.CancelVisitProposalRedirect(a.VisitRepo))
	})
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authorization(a.UserRepo))
		r.Get("/api/v1/visits", services.GetAllVisit(a.VisitRepo))
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
		r.Post("/api/v1/visits/confirmvisit", services.ConfirmVisitProposal(a.VisitRepo))
		r.Post("/api/v1/visits/cancelvisit", services.CancelVisitProposal(a.VisitRepo))
		r.Get("/api/v1/user/token", services.GetUserByToken(a.UserRepo))
		// Export Data

	})
	return r
}
