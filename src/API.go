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
	filter := r.URL.Query()
	// filter["Princess"] = true
	var res interface{}
	switch params["objects"] {
		// case "events":
		// 	res = IndexEvents(db, filter)
		case "gachas":
			res = IndexGachas(db, filter)
		// case "idols":
		// 	res = IndexIdols(db, filter)
		case "songs":
			res = IndexSongs(db, filter)
		case "cards":
			res = IndexCards(db, filter)
	}
	json.NewEncoder(w).Encode(res)
}

// func CreateObject(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
// 	checkError(err)
// 	params := mux.Vars(r)

// 	var res interface{}
// 	switch params["objects"] {
// 		// case "events":
// 		// 	res := CreateEvent(db, r)
// 		// 	json.NewEncoder(w).Encode(res)
// 		// case "gachas":
// 		// 	res := CreateGacha(db, r)
// 		// 	json.NewEncoder(w).Encode(res)
// 		// case "idols":
// 		// 	res := CreateIdol(db, r)
// 		// 	json.NewEncoder(w).Encode(res)
// 		case "songs":
// 			var song Song
// 			res = song.Create(db, r)
// 		// case "cards":
// 		// 	res := CreateCard(db, r)
// 		// 	json.NewEncoder(w).Encode(res)
// 	}
// 	json.NewEncoder(w).Encode(res)
// }

func ShowObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)

	var res interface{}
	switch params["objects"] {
		case "events":
			var event Event
			res = event.Show(db, params["id"])
		case "gachas":
			var gacha Gacha
			res = gacha.Show(db, params["id"])
		case "idols":
			var idol Idol
			res = idol.Show(db, params["id"])
		case "songs":
			var song Song
			res = song.Show(db, params["id"])
		case "cards":
			var card Card
			res = card.Show(db, params["id"])
	}
	json.NewEncoder(w).Encode(res)
}

// func UpdateObject(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
// 	checkError(err)
// 	params := mux.Vars(r)

// 	switch params["objects"] {
// 		case "events":
// 			res := UpdateEvent(db, r)
// 			json.NewEncoder(w).Encode(res)
// 		case "gachas":
// 			res := UpdateGacha(db, r)
// 			json.NewEncoder(w).Encode(res)
// 		case "idolss":
// 			res := UpdateIdol(db, r)
// 			json.NewEncoder(w).Encode(res)
// 		case "songs":
// 			res := UpdateSong(db, r)
// 			json.NewEncoder(w).Encode(res)
// 		case "cards":
// 			res := UpdateCard(db, r)
// 			json.NewEncoder(w).Encode(res)
// 	}

// }

// func DeleteObject(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
// 	checkError(err)
// 	params := mux.Vars(r)

// 	switch params["objects"] {
// 		case "events":
// 			res := DeleteEvent(db, params["id"])
// 			json.NewEncoder(w).Encode(res)
// 		case "gachas":
// 			res := DeleteGacha(db, params["id"])
// 			json.NewEncoder(w).Encode(res)
// 		case "idols":
// 			res := DeleteIdol(db, params["id"])
// 			json.NewEncoder(w).Encode(res)
// 		case "songs":
// 			res := DeleteSong(db, params["id"])
// 			json.NewEncoder(w).Encode(res)
// 		case "cards":
// 			res := DeleteCard(db, params["id"])
// 			json.NewEncoder(w).Encode(res)
// 	}
// }
