package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ANUMADHAV07/email-parser.git/internal/app"
	"github.com/ANUMADHAV07/email-parser.git/internal/routes"
	"github.com/joho/godotenv"
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

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file failed to load!")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientCallbackURL := os.Getenv("CLIENT_CALLBACK_URL")

	if clientID == "" || clientSecret == "" || clientCallbackURL == "" {
		log.Fatal("Environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
	}
}
