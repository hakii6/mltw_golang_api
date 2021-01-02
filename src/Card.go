package main

import (
	"encoding/json"
	// "time"
	"gorm.io/gorm"
	"net/http"
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

func IndexCards(db *gorm.DB) []Card {
	var cards []Card
	res := db.Select("id", "name_jp", "name_tw", "rarity", "total", "vocal", "dance", "visual", "limited", "date").Find(&cards)
	checkError(res.Error)
	return cards
}

func CreateCard(db *gorm.DB, r *http.Request) Card {
	var card Card
	err := json.NewDecoder(r.Body).Decode(&card)
	checkError(err)

	res := db.Select("id", "name_jp", "name_tw", "rarity", "limited", "date").Create(&card)
	checkError(res.Error)
	return card
}

func ShowCard(db *gorm.DB, id string) Card {
	var card Card
	res := db.Preload("Idol").Where("id = ?", id).First(&card)

	// res := db.Select("id", "name_jp", "name_tw", "rarity", "limited", "date").Where("id = ?", id).First(&card)
	checkError(res.Error)
	return card
}

func UpdateCard(db *gorm.DB, r *http.Request) Card {
	var card Card
	err := json.NewDecoder(r.Body).Decode(&card)
	checkError(err)

	res := db.Updates(&card)
	checkError(res.Error)
	return card
}

func DeleteCard(db *gorm.DB, id string) string {
	var card Card
	card.ID = id
	res := db.Delete(&card)
	checkError(res.Error)
	return "success"
}
