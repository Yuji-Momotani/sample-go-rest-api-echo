package main

import (
	"log"
	"sample-go-rest-api-echo/controller"
	"sample-go-rest-api-echo/db"
	"sample-go-rest-api-echo/repository"
	"sample-go-rest-api-echo/router"
	"sample-go-rest-api-echo/usecase"
	"sample-go-rest-api-echo/validator"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("NewDB in Error: %d\n", err)
	}

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
