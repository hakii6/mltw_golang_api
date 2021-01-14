package main

import (
	// "encoding/json"
	// "time"
	"gorm.io/gorm"
	"net/url"
	// "gorm.io/driver/mysql"

)

type Card struct {
	ID string `json:"ID"`
	Idol Idol
	IdolID string `json:"IdolID"`

	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	Rarity string `json:"Rarity"`
	Total int `json:"Total"`
	Vocal int `json:"Vocal"`
	Dance int `json:"Dance"`
	Visual int `json:"Visual"`
	Limited string `json:"Limited"`
	Date string `json:"Date"`

	ImageA string `json:"ImageA"`
	ImageB string `json:"ImageB"`

	GetCardID string `json:"GetCardID"`
	GetCardType string `json:"GetCardType"`
}

func IndexCards(db *gorm.DB, filter url.Values) []Card {
	var cards []Card
	temp := db.Select("id", "name_jp", "name_tw", "rarity", "total", "vocal", "dance", "visual", "limited", "date")
	if filter["type"] != nil {
		temp.Where("type = ?", filter["type"][0])
	}
	if filter["rarity"] != nil {
		temp.Where("rarity = ?", filter["rarity"][0])
	}
	if filter["year"] != nil {
		temp.Where("date BETWEEN ? AND ?", filter["year"][0] + "-01-01", filter["year"][0] + "-12-31")
	}
	res := temp.Find(&cards)
	checkError(res.Error)
	return cards
}


func (card *Card) Show(db *gorm.DB, id string) *Card{
	res := db.Preload("Idol").Where("id = ?", id).First(&card)
	checkError(res.Error)

	return card
}