package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/brandon-hiles/medium/golang-auth/pkg/db"
	"github.com/joho/godotenv"
)

type application struct {
	logger *log.Logger
	pg     db.DatabaseInterface
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

	// Initiate Server
	app.logger.Println("Starting server on port", port)
	server := http.ListenAndServe(":"+strconv.Itoa(port), app.routes())
	if server != nil {
		app.logger.Fatal("Error occurred")
	}
}
