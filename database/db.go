package database

import (
		"log"
		"github.com/glebarez/sqlite"
		"gorm.io/gorm"
	
)

var DB *gorm.DB

func Data(){
	// form sqlite database
	db, err := gorm.Open(
		sqlite.Open("checker.db"),
		&gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}

	DB = db

}
