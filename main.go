package main

import (
	"todo-rest-api-3/controller"
	"todo-rest-api-3/db"
	"todo-rest-api-3/repository"
	"todo-rest-api-3/router"
	"todo-rest-api-3/usecase"
	"todo-rest-api-3/validator"
)

func main() {
	db := db.NewDB()
	//処理
	//repository
	//usecase
	//controller
	//router → echo.Echo
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserUsecase(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
