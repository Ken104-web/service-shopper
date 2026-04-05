package main

import (
	"log"
	"math/rand"
	"time"

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
	db.Exec("DELETE FROM customers")
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
	db.Exec("DELETE FROM services")

	services := []model.Service{}
		// append that data
		fake = faker.New()
		for i:=0; i<3; i++{
			s := fake.Person()
			rating := fake.IntBetween(0, 10)
			servant := model.Service{
				Name: s.Name(),
				Rating: rating,
			}
			services = append(services, servant)
		}
		// save to DB 
		res :=db.Create(&services)
		if res.Error != nil{
			log.Fatalf("customer cannot be served due to lack of servant: %v", res.Error)
		}
	
	db.Exec("DELETE FROM products")	
	
	products := []model.Product{}
	//  list the array data
	items := [...]string{"Bread", "Milk", "Eggs", "Fresh fruits"}
	fake = faker.New()	
	for i:=0; i<3; i++{
		// to avoid repeate of the same random object
		rand.NewSource(time.Now().UnixNano())

		product_name := items[rand.Intn(len(items))] 
		price := fake.IntBetween(8, 50)
		// lets picks randomly here
		customer := customers[rand.Intn(len(customers))]
		service := services[rand.Intn(len(services))]
		
		// call the here
		produce := model.Product{
			Item_Name: product_name,	
			Price: price,
			CustomerID: customer.ID,
			ServiceID: service.ID,

		}
		products = append(products, produce)
	}
	result := db.Create(&products)
	if result.Error != nil{
		log.Fatalf("pfft what more can i say its not there: %v",  result.Error)
	}
	
	// fill the join table 
	db.Exec("DELETE FROM customers")	
	for i := range customers{
		customer := &customers[i]

		// fill random services
		for i := 0; i<3; i++ {
			service := services[rand.Intn(len(services))]

			err := db.Model(customer).Association("Services").Append(&service)
				
			if err != nil {
				log.Fatal(err)
				}
		}
		
	}
	err = db.AutoMigrate(&model.Customer{}, &model.Service{}, &model.Product{})
if err != nil{
	log.Fatal(err)
	}

}
