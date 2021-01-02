package main

import (
	"encoding/json"
	"time"
	"gorm.io/gorm"
	"net/http"
	// "gorm.io/driver/mysql"

)

type Song struct {
	ID string `json:"ID"`
	NameJP string `json:"NameJP"`
	NameTW string `json:"NameTW"`
	BPM string `json:"BPM"`
	Length string `json:"Length"`
	Date time.Time `json:"Date"`
	Image string `json:"Image"`
	Type string `json:"Type"`

	EzLv int `json:"EzLv"`
	NmLv int `json:"NmLv"`
	HrLv int `json:"HrLv"`
	Hr2Lv int `json:"Hr2Lv"`
	ExLv int `json:"ExLv"`

	EzNotes int `json:"EzNotes"`
	NmNotes int `json:"NmNotes"`
	HrNotes int `json:"HrNotes"`
	Hr2Notes int `json:"Hr2Notes"`
	ExNotes int `json:"ExNotes"`
}

func IndexSongs(db *gorm.DB) []Song {
	var songs []Song
	res := db.Select("id", "name_jp", "name_tw", "BPM", "length", "image", "date", "type").Find(&songs)
	checkError(res.Error)
	return songs
}

func CreateSong(db *gorm.DB, r *http.Request) Song {
	var song Song
	err := json.NewDecoder(r.Body).Decode(&song)
	checkError(err)

	res := db.Create(&song)
	checkError(res.Error)
	return song
}

func ShowSong(db *gorm.DB, id string) Song {
	var song Song
	res := db.Where("id = ?", id).First(&song)
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
