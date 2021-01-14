package main

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	// "gorm.io/gorm"
	// "encoding/json"
	// "gorm.io/driver/mysql"
)

func SetHeader(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
	    next.ServeHTTP(w, r)
    })
}

func main() {


	// Init Router
	Router := mux.NewRouter()

	// V1 
	RV1 := Router.PathPrefix("/v1").Subrouter()
	
	// Index & CRUD
	RV1.HandleFunc("/{objects}", IndexObjects).Methods("GET")
	// RV1.HandleFunc("/{objects}", CreateObject).Methods("POST")
	RV1.HandleFunc("/{objects}/{id:[0-9]+}", ShowObject).Methods("GET")
	// RV1.HandleFunc("/{objects}/{id}", UpdateObject).Methods("PATCH")
	// RV1.HandleFunc("/{objects}/{id}", DeleteObject).Methods("DELETE")

	RV1.Use(SetHeader)



    log.Fatal(http.ListenAndServe(":8002", Router))

}