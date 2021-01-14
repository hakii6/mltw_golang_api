package main

import (
	// "encoding/json"
	"time"
	"gorm.io/gorm"
	"net/url"
	// "net/http"
	// "gorm.io/driver/mysql"

)

type Gacha struct {
	ID string `json:"ID"`
	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	StartDate time.Time `json:"StartDate"`
	EndDate time.Time `json:"EndDate"`
	Ori_url string `json:"Ori_url"`
	Image string `json:"Image"`
	Cards []Card `gorm:"polymorphic:GetCard;polymorphicValue:Gacha"`
}

func IndexGachas(db *gorm.DB, filter url.Values) []Gacha {
	var gachas []Gacha
	temp := db.Preload("Cards")
	if filter["year"] != nil {
		temp.Where("(start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?)", 
			filter["year"][0] + "-01-01", filter["year"][0] + "-12-31", filter["year"][0] + "-01-01", filter["year"][0] + "-12-31")
	}
	res := temp.Find(&gachas)
	checkError(res.Error)
	return gachas
}

func (gacha *Gacha) Show(db *gorm.DB, id string) *Gacha{
	res := db.Select("id", "name_jp", "name_tw", "start_date", "end_date", "ori_url").Where("id = ?", id).First(&gacha)
	checkError(res.Error)

	return gacha
}
