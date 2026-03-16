package main

import (
	"log"

	"service-shopper/model"
	
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
func main(){
	// form sqlite database
	db, err := gorm.Open(
		sqlite.Open("checker.db"),
		&gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.Customer{}, &model.Service{}, &model.Product{})
	if err != nil{
		log.Fatal(err)
	}
	
}
