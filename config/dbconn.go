package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DBinit(localConfig *LocalConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", localConfig.DBUser, localConfig.DBPassword, localConfig.Host, localConfig.DBPort, localConfig.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Connection established successfully")
	return db, nil
	//connStr := "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
}
