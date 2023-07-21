package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/brandon-hiles/medium/golang-auth/pkg/auth"
	"github.com/brandon-hiles/medium/golang-auth/pkg/email"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // one megabyte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

// Write out a JSON Response
func writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return writeJSON(w, statusCode, payload)
}

func (app *application) helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

// Create a new user handler
func (app *application) create_user(w http.ResponseWriter, r *http.Request) {
	app.logger.Printf("%s /create_user", r.Method)
	var requestPayload struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}

	err := readJSON(w, r, &requestPayload)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.pg.InsertUser(requestPayload.FirstName, requestPayload.LastName, requestPayload.PhoneNumber, requestPayload.Email, requestPayload.Password)
	if err != nil {
		errorJSON(w, errors.New("user already exists"), http.StatusBadRequest)
		app.logger.Print("User already exists")
		return
	}
}

func (app *application) send_signup_email(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		Email string `json:"email"`
	}
	err := readJSON(w, r, &requestPayload)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	to := []string{requestPayload.Email}
	url := "http://localhost:9000/login"

	// validate user against database
	user, err := app.pg.GetUserByEmail(requestPayload.Email)
	if err != nil {
		errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	email.SendSignupVerificationEmail(app.email, to, user.Email, url)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json_error := readJSON(w, r, &requestPayload)
	if json_error != nil {
		errorJSON(w, json_error, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.pg.GetUserByEmail(requestPayload.Email)
	if err != nil {
		errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := auth.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.logger.Println(err)
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	writeJSON(w, http.StatusAccepted, tokens)
}

func (app *application) refresh(w http.ResponseWriter, r *http.Request) {

}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {

}
