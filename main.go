package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jonathanmorais/api-spotify-temp/routes"
	"log"
	"net/http"
)

// main Search
func main() {
   fmt.Print("hello")
   r := mux.NewRouter()
   r.HandleFunc("/", routes.HomeHandler).Methods("GET")
   r.HandleFunc("/receive", routes.ReceiveCoordinates).Methods("POST")
   http.Handle("/", r)
   log.Fatal(http.ListenAndServe(":8090", r))
}