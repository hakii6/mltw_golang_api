package main

import (
	"encoding/json"
	"time"
	"gorm.io/gorm"
	"net/http"
	// "gorm.io/driver/mysql"

)

type Gacha struct {
	ID string `json:"id"`
	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	StartDate time.Time `json:"StartDate"`
	EndDate time.Time `json:"EndDate"`
	Ori_url string `json:"Ori_url"`
}

func IndexGachas(db *gorm.DB) []Gacha {
	var gachas []Gacha
	res := db.Select("id", "name_jp", "name_tw", "start_date", "end_date", "ori_url").Find(&gachas)
	checkError(res.Error)
	return gachas
}

func CreateGacha(db *gorm.DB, r *http.Request) Gacha {
	var gacha Gacha
	err := json.NewDecoder(r.Body).Decode(&gacha)
	checkError(err)

	res := db.Select("id", "name_jp", "name_tw", "start_date", "end_date", "ori_url").Create(&gacha)
	checkError(res.Error)
	return gacha
}

func ShowGacha(db *gorm.DB, id string) Gacha {
	var gacha Gacha
	res := db.Select("id", "name_jp", "name_tw", "start_date", "end_date", "ori_url").Where("id = ?", id).First(&gacha)
	checkError(res.Error)
	return gacha
}

func UpdateGacha(db *gorm.DB, r *http.Request) Gacha {
	var gacha Gacha
	err := json.NewDecoder(r.Body).Decode(&gacha)
	checkError(err)

	res := db.Updates(&gacha)
	checkError(res.Error)
	return gacha
}

func DeleteGacha(db *gorm.DB, id string) string {
	var gacha Gacha
	gacha.ID = id
	res := db.Delete(&gacha)
	checkError(res.Error)
	return "success"
}
