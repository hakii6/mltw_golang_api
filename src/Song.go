package main

import (
	// "encoding/json"
	// "log"
	"time"
	"gorm.io/gorm"
	"net/url"
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

// func (s Song) Index(db *gorm.DB) []Song {
// 	var songs []Song
// 	res := db.Select("id", "name_jp", "name_tw", "BPM", "length", "image", "date", "type").Find(&songs)
// 	checkError(res.Error)
// 	return songs
// }

func IndexSongs(db *gorm.DB, filter url.Values) []Song {
	var songs []Song
	temp := db.Select("id", "name_jp", "name_tw", "BPM", "length", "image", "date", "type")
	if filter["type"] != nil {
		temp.Where("type = ?", filter["type"][0])
	}
	if filter["year"] != nil {
		temp.Where("date BETWEEN ? AND ?", filter["year"][0] + "-01-01", filter["year"][0] + "-12-31")
	}
	// for key, _ := range filter {
	// 	temp.Where("type = ?", key)
	// }
	res := temp.Find(&songs)
	checkError(res.Error)

	return songs
}

func (song *Song) Show(db *gorm.DB, id string) *Song{
	res := db.Where("id = ?", id).First(&song)
	checkError(res.Error)

	return song
}

// func (song *Song) Create(db *gorm.DB, r *http.Request) *Song{
// 	err := json.NewDecoder(r.Body).Decode(&song)
// 	checkError(err)

// 	res := db.Create(&song)
// 	checkError(res.Error)

// 	return song
// }

// func (song *Song) Update(db *gorm.DB, r *http.Request) *Song {
// 	err := json.NewDecoder(r.Body).Decode(&song)
// 	checkError(err)

// 	res := db.Updates(&song)
// 	checkError(res.Error)

// 	return song
// }

// func (song *Song) Delete(db *gorm.DB, id string) string {
// 	song.ID = id
// 	res := db.Delete(&song)
// 	checkError(res.Error)

// 	return "success"
// }
