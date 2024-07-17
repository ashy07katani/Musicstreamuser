package main

import "log"

func main() {
	db, err := DBinit()
	if err != nil {
		log.Fatalf("Error initializing connection with Database %v", err)
	}
	log.Printf("%+v", db)

}
