package main

import (
	"encoding/json"
	"time"
	"gorm.io/gorm"
	"net/http"
	// "gorm.io/driver/mysql"

)

type Song struct {
	ID string `json:"id"`
	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	BPM string `json:"BPM"`
	Length string `json:"Length"`
	Date time.Time `json:"Date"`
}

func IndexSongs(db *gorm.DB) []Song {
	var songs []Song
	res := db.Select("id", "name_jp", "name_tw", "BPM", "length", "date").Find(&songs)
	checkError(res.Error)
	return songs
}

func CreateSong(db *gorm.DB, r *http.Request) Song {
	var song Song
	err := json.NewDecoder(r.Body).Decode(&song)
	checkError(err)

	res := db.Select("id", "name_jp", "name_tw", "BPM", "length", "date").Create(&song)
	checkError(res.Error)
	return song
}

func ShowSong(db *gorm.DB, id string) Song {
	var song Song
	res := db.Select("id", "name_jp", "name_tw", "BPM", "length", "date").Where("id = ?", id).First(&song)
	checkError(res.Error)
	return song
}

func UpdateSong(db *gorm.DB, r *http.Request) Song {
	var song Song
	err := json.NewDecoder(r.Body).Decode(&song)
	checkError(err)

	res := db.Updates(&song)
	checkError(res.Error)
	return song
}

func DeleteSong(db *gorm.DB, id string) string {
	var song Song
	song.ID = id
	res := db.Delete(&song)
	checkError(res.Error)
	return "success"
}
