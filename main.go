package main

import (
	"log"

	"service-shopper/model"

	"github.com/glebarez/sqlite"
	"github.com/jaswdr/faker/v2"
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
	// insert customer data into tables 
	// define the slice(literal)
	customers := []model.Customer{}
	// append them to the slice 
		fake := faker.New()
		for i:=0; i<3; i++{
			c := fake.Person()
			customer := model.Customer{
				Name: c.Name(),
			}
			customers = append(customers, customer)
		}
		// save to the DB 
		r :=db.Create(&customers)
		if r.Error != nil{
			log.Fatalf("no record of customer: %v", r.Error)
		}


	err = db.AutoMigrate(&model.Customer{}, &model.Service{}, &model.Product{})
	if err != nil{
		log.Fatal(err)
	}
	
}
