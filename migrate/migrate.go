package main

import (
	"fmt"
	"sample-go-rest-api-echo/db"
	"sample-go-rest-api-echo/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
