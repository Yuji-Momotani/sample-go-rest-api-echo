package main

import (
	"fmt"
	"log"
	"sample-go-rest-api-echo/db"
	"sample-go-rest-api-echo/model"
)

func main() {
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalf("NewDB in Error: %d\n", err)
	}
	defer fmt.Println("Successfully Migrate")
	defer func() {
		err := db.CloseDB(dbConn)
		if err != nil {
			log.Fatalf("CloseDB in Error: %d\n", err)
		}
	}()
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
