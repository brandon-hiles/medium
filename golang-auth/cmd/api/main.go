package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

type application struct {
	logger *log.Logger
}

func main() {
	var app application
	app.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	port := 8001
	// Initiate Server
	app.logger.Println("Starting server on port", port)
	server := http.ListenAndServe(":"+strconv.Itoa(port), app.routes())
	if server != nil {
		log.Fatal("Error occurred")
	}
}
