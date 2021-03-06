package main

import (
	// "encoding/json"
	// "time"
	"gorm.io/gorm"
	"net/url"
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

func IndexIdols(db *gorm.DB, filter url.Values) []Idol {
	var idols []Idol
	res := db.Select("id", "name_jp", "name_tw", "type", "thumbnail", "intro").Find(&idols)
	checkError(res.Error)
	return idols
}

func (idol *Idol) Show(db *gorm.DB, id string) *Idol{
	res := db.Preload("Songs").Preload("Cards").Where("id = ?", id).First(&idol)
	checkError(res.Error)

	return idol
}