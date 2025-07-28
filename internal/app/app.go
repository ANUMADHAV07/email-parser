package app

import (
	"fmt"
	"net/http"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
