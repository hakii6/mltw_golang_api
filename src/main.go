package main

import (
	// "encoding/json"
	_ "fmt"
	"net/http"
	"log"
	"os"
	"github.com/gorilla/mux"
	"github.com/kardianos/service"
)
var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	// Do work here

	// Init Router
	Router := mux.NewRouter()



	// Index & CRUD
	Router.HandleFunc("/api/v0/{objects}", IndexObjects).Methods("GET")
	// Router.HandleFunc("/api/v0/{objects}", CreateObject).Methods("POST")
	Router.HandleFunc("/api/v0/{objects}/{id}", ShowObject).Methods("GET")
	// Router.HandleFunc("/api/v0/{objects}/{id}", UpdateObject).Methods("PATCH")
	// Router.HandleFunc("/api/v0/{objects}/{id}", DeleteObject).Methods("DELETE")

    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    //     fmt.Fprintf(w, "hiii")
    // })

    log.Fatal(http.ListenAndServe(":8001", Router))

}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoAPI",
		DisplayName: "Go API",
		Description: "Mltw Golang API",
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// if in windows install and start uself
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if os.Args[1] != "start" {
			return
		}
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}


}