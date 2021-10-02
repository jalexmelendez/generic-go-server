package main

import (
	"fmt"
	controller "go/server/controllers"
	"log"
	"net/http"
)

/*Go Server*/
func runserver() {
	http.HandleFunc("/user", controller.UserHandler)
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

/*Server App Entry Point, You can add functions to run processes*/
func main() {
	runserver()
}
