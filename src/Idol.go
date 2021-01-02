package main

import (
	"encoding/json"
	// "time"
	"gorm.io/gorm"
	"net/http"
	// "gorm.io/driver/mysql"

)

type Idol struct {
	ID string `json:"ID"`
	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	Type string `json:"Type"`
	Thumbnail string `json:"Thumbnail"`
	Intro string `json:"Intro"`
	Songs []Song `gorm:"many2many:idol_song"`
	Cards []Card `gorm:"foreignKey:IdolID"`
}

func IndexIdols(db *gorm.DB) []Idol {
	var idols []Idol
	res := db.Select("id", "name_jp", "name_tw", "type", "thumbnail", "intro").Find(&idols)
	checkError(res.Error)
	return idols
}

func CreateIdol(db *gorm.DB, r *http.Request) Idol {
	var idol Idol
	err := json.NewDecoder(r.Body).Decode(&idol)
	checkError(err)

	res := db.Select("id", "name_jp", "name_tw", "type", "thumbnail", "intro").Create(&idol)
	checkError(res.Error)
	return idol
}

func ShowIdol(db *gorm.DB, id string) Idol {
	var idol Idol
	db.Preload("Songs").Preload("Cards").Where("id = ?", id).First(&idol)
		// res := db.Select("id", "name_jp", "name_tw", "type", "thumbnail", "intro").Where("id = ?", id).First(&idol)

	// res := db.Select("id", "name_jp", "name_tw", "type", "thumbnail", "intro").Where("id = ?", id).First(&idol)
	// checkError(res.Error)
	return idol
}

func UpdateIdol(db *gorm.DB, r *http.Request) Idol {
	var idol Idol
	err := json.NewDecoder(r.Body).Decode(&idol)
	checkError(err)

	res := db.Updates(&idol)
	checkError(res.Error)
	return idol
}

func DeleteIdol(db *gorm.DB, id string) string {
	var idol Idol
	idol.ID = id
	res := db.Delete(&idol)
	checkError(res.Error)
	return "success"
}
