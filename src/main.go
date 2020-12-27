package main

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {

	// Init Router
	Router := mux.NewRouter()



	// Index & CRUD
	Router.HandleFunc("/api/v0/{objects}", IndexObjects).Methods("GET")
	Router.HandleFunc("/api/v0/{objects}", CreateObject).Methods("POST")
	Router.HandleFunc("/api/v0/{objects}/{id}", ShowObject).Methods("GET")
	Router.HandleFunc("/api/v0/{objects}/{id}", UpdateObject).Methods("POST")
	Router.HandleFunc("/api/v0/{objects}/{id}", DeleteObject).Methods("Delete")

    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "hiii")
    // })

    log.Fatal(http.ListenAndServe(":8001", Router))
}