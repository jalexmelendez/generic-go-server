package controller

import (
	"fmt"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It Works.")
}
