package main

import (
	"fmt"
	"net/http"
)

func (app *application) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func (app *application) create_user(w http.ResponseWriter, r *http.Request) {

}

func (app *application) send_signup_email(w http.ResponseWriter, r *http.Request) {

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {

}

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {

}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {

}
