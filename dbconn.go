package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type LocalConfig struct {
	Host       string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	Port       int    `env:"DB_PORT"`
}

func GetDBConfig() (*LocalConfig, error) {
	var localConfig LocalConfig
	//will use godotenv to load the file
	err := godotenv.Load("local.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	//from the loaded file we will parse the environment variable into localconfig instance
	if err := env.Parse(&localConfig); err != nil {
		log.Fatalf("Error reading the environment variable %v", err)
		return nil, err
	}
	log.Printf("%+v\n", localConfig)
	return &localConfig, nil
}

func DBinit() (*sql.DB, error) {
	localConfig, err := GetDBConfig()
	if err != nil {
		log.Fatalf("Error initializing database configurations %v", err)
		return nil, err
	}
	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", localConfig.DBUser, localConfig.DBPassword, localConfig.Host, localConfig.DBPort, localConfig.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Connection established successfully")
	defer db.Close()
	return db, nil
	//connStr := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
}
