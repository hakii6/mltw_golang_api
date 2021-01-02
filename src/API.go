package main

import (
	"encoding/json"
	"net/http"
	// "fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)


func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func IndexObjects(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	switch params["objects"] {
		case "events":
			res := IndexEvents(db)
			json.NewEncoder(w).Encode(res)
		case "gachas":
			res := IndexGachas(db)
			json.NewEncoder(w).Encode(res)
		case "idols":
			res := IndexIdols(db)
			json.NewEncoder(w).Encode(res)		
		case "songs":
			res := IndexSongs(db)
			json.NewEncoder(w).Encode(res)
		case "cards":
			res := IndexCards(db)
			json.NewEncoder(w).Encode(res)	
	}

	
}

func CreateObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	switch params["objects"] {
		case "events":
			res := CreateEvent(db, r)
			json.NewEncoder(w).Encode(res)
		case "gachas":
			res := CreateGacha(db, r)
			json.NewEncoder(w).Encode(res)
		case "idols":
			res := CreateIdol(db, r)
			json.NewEncoder(w).Encode(res)
		case "songs":
			res := CreateSong(db, r)
			json.NewEncoder(w).Encode(res)
		case "cards":
			res := CreateCard(db, r)
			json.NewEncoder(w).Encode(res)
	}
}

func ShowObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	switch params["objects"] {
		case "events":
			res := ShowEvent(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "gachas":
			res := ShowGacha(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "idols":
			res := ShowIdol(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "songs":
			res := ShowSong(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "cards":
			res := ShowCard(db, params["id"])
			json.NewEncoder(w).Encode(res)
	}
}

func UpdateObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	switch params["objects"] {
		case "events":
			res := UpdateEvent(db, r)
			json.NewEncoder(w).Encode(res)
		case "gachas":
			res := UpdateGacha(db, r)
			json.NewEncoder(w).Encode(res)
		case "idolss":
			res := UpdateIdol(db, r)
			json.NewEncoder(w).Encode(res)
		case "songs":
			res := UpdateSong(db, r)
			json.NewEncoder(w).Encode(res)
		case "cards":
			res := UpdateCard(db, r)
			json.NewEncoder(w).Encode(res)
	}

}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	switch params["objects"] {
		case "events":
			res := DeleteEvent(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "gachas":
			res := DeleteGacha(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "idols":
			res := DeleteIdol(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "songs":
			res := DeleteSong(db, params["id"])
			json.NewEncoder(w).Encode(res)
		case "cards":
			res := DeleteCard(db, params["id"])
			json.NewEncoder(w).Encode(res)
	}
}
