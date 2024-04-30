package main

import (
	"log"
	"main/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", auth.Auth_check).
		Methods("GET", "POST")

	route.HandleFunc("/signin", auth.Signin)

	route.HandleFunc("/welcome", auth.Welcome)

	server := &http.Server{
		Handler: route,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(server.ListenAndServe())
}
