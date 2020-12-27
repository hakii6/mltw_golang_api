package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

const (
    host     = "localhost"
    database = "test2"
    user     = "root"
    password = "zxcv2587"
    connect = user + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8&parseTime=true"
)

// func (song *Song) Show(db *gorm.DB, id string) string {
// 	res := db.Select("id", "name_jp", "name_tw", "BPM", "length", "date").Where("id = ?", id).First(&song)
// 	checkError(res.Error)
// 	return "Success"
// }  

// type Card struct {
	
// }

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func IndexObjects(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)

	switch params["objects"] {
		case "songs":
			res := IndexSongs(db)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		case "cards":
			res := IndexCards(db)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		}
}

func CreateObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)

	switch params["objects"] {
		case "songs":
			res := CreateSong(db, r)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		case "cards":
			fmt.Println("cards")
	}
}

func ShowObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)

	switch params["objects"] {
		case "songs":
			res := ShowSong(db, params["id"])
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		case "cards":
			fmt.Println("cards")
	}
}

func UpdateObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)

	switch params["objects"] {
		case "songs":
			res := UpdateSong(db, r)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		case "cards":
			fmt.Println("cards")
	}

}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	checkError(err)
	params := mux.Vars(r)

	switch params["objects"] {
		case "songs":
			res := DeleteSong(db, params["id"])
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
		case "cards":
			fmt.Println("cards")
	}
}
