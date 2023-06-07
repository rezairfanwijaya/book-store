package main

import (
	"book-store/database"
	"log"
)

func main() {
	// connection
	_, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Success connection")
}
