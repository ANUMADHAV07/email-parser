package app

import "github.com/ANUMADHAV07/email-parser.git/internal/handlers"

type Application struct {
	Handler *handlers.Handler
}

func NewApplication() *Application {
	handler := handlers.NewHandler()
	return &Application{
		Handler: handler,
	}
}
