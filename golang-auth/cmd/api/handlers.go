package main

import (
	"fmt"
	"html"
	"net/http"
)

func (app *application) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
