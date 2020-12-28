package main

import (
	"encoding/json"
	"time"
	"gorm.io/gorm"
	"net/http"
	// "gorm.io/driver/mysql"

)

type Event struct {
	ID string `json:"id"`
	Name_jp string `json:"NameJP"`
	Name_tw string `json:"NameTW"`
	StartDate time.Time `json:"StartDate"`
	BoostDate time.Time `json:"BoostDate"`
	EndDate time.Time `json:"EndDate"`
	Ori_url string `json:"Ori_url"`
	Type string `json:"Type"`
}

func IndexEvents(db *gorm.DB) []Event {
	var events []Event
	res := db.Select("id", "name_jp", "name_tw", "start_date", "boost_date", "end_date", "ori_url", "type").Find(&events)
	checkError(res.Error)
	return events
}

func CreateEvent(db *gorm.DB, r *http.Request) Event {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	checkError(err)

	res := db.Select("id", "name_jp", "name_tw", "start_date", "boost_date", "end_date", "ori_url", "type").Create(&event)
	checkError(res.Error)
	return event
}

func ShowEvent(db *gorm.DB, id string) Event {
	var event Event
	res := db.Select("id", "name_jp", "name_tw", "start_date", "boost_date", "end_date", "ori_url", "type").Where("id = ?", id).First(&event)
	checkError(res.Error)
	return event
}

func UpdateEvent(db *gorm.DB, r *http.Request) Event {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	checkError(err)

	res := db.Updates(&event)
	checkError(res.Error)
	return event
}

func DeleteEvent(db *gorm.DB, id string) string {
	var event Event
	event.ID = id
	res := db.Delete(&event)
	checkError(res.Error)
	return "success"
}
