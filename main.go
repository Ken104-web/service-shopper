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
	service := []model.Service{}
	// append that data 
	fake := faker.New()
	for i:=0; i<3; i++{
		s := fake.Person()
		// r := fake.RandomDigit()
		servant := model.Service{
			Name: s.Name()
			// Rating: r.Rating()
		}
		servants =append(servants, servant)
	}
	// save to DB 
	r :=db.Create(&service)
	if r.Error != nil{
		log.Fatal("customer cannot be served due to lack of servant: %v", r.Error)
	}
	
	product := []model.Product{}
	// list the array data
	items := [...]string{"Bread", "Milk", "Eggs", "Fresh fruits"}

	err = db.AutoMigrate(&model.Customer{}, &model.Service{}, &model.Product{})
	if err != nil{
		log.Fatal(err)
	}
	
}
