package main

import (
	"Musicstreamuser/config"
	"Musicstreamuser/controller"
	"Musicstreamuser/repository"
	"Musicstreamuser/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	localConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error initializing database configurations %v", err)
		return
	}
	db, err := DBinit(localConfig)
	if err != nil {
		log.Fatalf("Error initializing connection with Database %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error keeping connection with Database %v", err)
	}
	log.Printf("%+v", db)
	repository := repository.NewRepository(db)
	userService := service.NewService(repository)
	userController := controller.NewController(userService)
	router := mux.NewRouter()
	userController.RegisterRoutes(router)
	log.Println("Server is about to start")
	http.ListenAndServe(fmt.Sprintf(":%v", localConfig.Port), router)
	defer db.Close()

}
