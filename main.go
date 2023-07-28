package main

import (
	"golang_clean_architecture_sample/controller"
	"golang_clean_architecture_sample/db"
	"golang_clean_architecture_sample/repository"
	"golang_clean_architecture_sample/router"
	"golang_clean_architecture_sample/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
