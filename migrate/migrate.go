package main

import (
	"fmt"
	"todo-rest-api-3/db"
	"todo-rest-api-3/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
