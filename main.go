package main

import (
	"log"
	"todo-rest-api-3/controller"
	"todo-rest-api-3/db"
	"todo-rest-api-3/repository"
	"todo-rest-api-3/router"
	"todo-rest-api-3/usecase"
	"todo-rest-api-3/validator"
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
