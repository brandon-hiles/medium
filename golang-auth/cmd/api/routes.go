package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/hello_world/", app.helloWorld)

	// Authentication routes
	r.HandleFunc("/create_user", app.create_user)
	r.HandleFunc("/send_signup_email", app.send_signup_email)
	r.HandleFunc("/authenticate", app.authenticate)
	r.HandleFunc("/refresh", app.refresh)
	r.HandleFunc("/logout", app.logout)

	http.Handle("/", r)
	return r
}
