package routes

import (
	"github.com/ANUMADHAV07/email-parser.git/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", app.HealthCheck)

	return r
}
