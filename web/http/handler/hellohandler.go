package handler

import (
	"fmt"
	"net/http"
)

// Hello world router
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
