package main

import (
	"Musicstreamuser/config"
	"Musicstreamuser/controller"
	"Musicstreamuser/repository"
	"Musicstreamuser/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	localConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error initializing database configurations %v", err)
		return
	}
	db, err := config.DBinit(localConfig)
	if err != nil {
		log.Fatalf("Error initializing connection with Database %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error keeping connection with Database %v", err)
	}
	log.Printf("%+v", db)
	repository := repository.NewRepository(db)
	userService := service.NewService(repository, localConfig.JwtSecret)
	userController := controller.NewController(userService)
	router := mux.NewRouter()
	userController.RegisterRoutes(router)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", localConfig.Port),
		Handler: router,
	}
	log.Println("Server is about to start")
	go func() {
		if err = server.ListenAndServe(); err != nil {
			log.Fatalf("Error in starting the server: %v", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	sig := <-sigChan
	log.Fatalf("Gracefully Shutting down due to the signal received: %v", sig)
	shutDownCtx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	err = server.Shutdown(shutDownCtx)
	defer cancel()
	defer db.Close()

}
