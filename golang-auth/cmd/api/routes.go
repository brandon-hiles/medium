package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/hello_world/", app.HelloWorld)
	http.Handle("/", r)
	return r
}
