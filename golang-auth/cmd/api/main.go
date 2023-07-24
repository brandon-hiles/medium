package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/brandon-hiles/medium/golang-auth/pkg/auth"
	"github.com/brandon-hiles/medium/golang-auth/pkg/db"
	"github.com/joho/godotenv"
)

type application struct {
	logger *log.Logger
	pg     db.DatabaseInterface
	auth   auth.Auth
}

// Basic way to load in a environment variables from .env file
func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	var app application
	app.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	port := 8001

	// Connect to our postgres database
	pg_user := goDotEnvVariable("DB_USER")
	pg_host := goDotEnvVariable("DB_HOST")
	pg_name := goDotEnvVariable("DB_NAME")
	pg_password := goDotEnvVariable("DB_PASSWORD")

	connection := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pg_user, pg_password, pg_host, pg_name)

	database, err := db.OpenDB(connection)
	if err != nil {
		app.logger.Fatal(err)
	}
	app.pg = &db.PostgresConnection{DB: database}
	defer database.Close()

	// Setup our Auth object
	// Setup our JWT authentication service
	app.auth = auth.Auth{
		Issuer:        goDotEnvVariable("JWT_ISSUER"),
		Audience:      goDotEnvVariable("JWT_AUDIENCE"),
		Secret:        goDotEnvVariable("JWT_SECRET"),
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "refresh_token",
		CookieDomain:  goDotEnvVariable("JWT_COOKIE_DOMAIN"),
	}

	// Initiate Server
	app.logger.Println("Starting server on port", port)
	server := http.ListenAndServe(":"+strconv.Itoa(port), app.routes())
	if server != nil {
		app.logger.Fatal("Error occurred")
	}
}
