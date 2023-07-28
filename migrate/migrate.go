package main

import (
	"fmt"
	"golang_clean_architecture_sample/db"
	"golang_clean_architecture_sample/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
