package main

import (
	"fmt"
	"net/http"

	"github.com/ANUMADHAV07/email-parser.git/internal/app"
	"github.com/ANUMADHAV07/email-parser.git/internal/routes"
)

func main() {
	app := app.NewApplication()

	r := routes.SetupRoutes(app)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}

	fmt.Println("server started")
}
