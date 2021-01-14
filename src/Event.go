package main

import (
	// "encoding/json"
	"time"
	"gorm.io/gorm"
	"net/url"
	// "gorm.io/driver/mysql"

)

type Event struct {
	ID string `json:"ID"`
	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	StartDate time.Time `json:"StartDate"`
	BoostDate time.Time `json:"BoostDate"`
	EndDate time.Time `json:"EndDate"`
	Ori_url string `json:"Ori_url"`
	Type string `json:"Type"`
	Image string `json:"Image"`
	Cards []Card `gorm:"polymorphic:GetCard;polymorphicValue:Event"`
}

func IndexEvents(db *gorm.DB, filter url.Values) []Event {
	var events []Event
	temp := db.Preload("Cards")
	if filter["year"] != nil {
		temp.Where("(start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?)", 
			filter["year"][0] + "-01-01", filter["year"][0] + "-12-31", filter["year"][0] + "-01-01", filter["year"][0] + "-12-31")
	}
	res := temp.Find(&events)
	checkError(res.Error)
	return events
}

func (event *Event) Show(db *gorm.DB, id string) *Event {
	res := db.Preload("Cards").Where("id = ?", id).First(&event)
	checkError(res.Error)

	return event
}
