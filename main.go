package main

import (
	"log"
	"math/rand"
	"time"

	"service-shopper/model"
	"service-shopper/database"

	"github.com/jaswdr/faker/v2"
)

func main(){
	//call the database func
	database.Data()
	// insert customer data into tables 
	// define the slice(literal)

	database.DB.Exec("DELETE FROM customers")
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
		r :=database.DB.Create(&customers)
		if r.Error != nil{
			log.Fatalf("no record of customer: %v", r.Error)
		}
	database.DB.Exec("DELETE FROM services")

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
		res :=database.DB.Create(&services)
		if res.Error != nil{
			log.Fatalf("customer cannot be served due to lack of servant: %v", res.Error)
		}
	
	database.DB.Exec("DELETE FROM products")	
	
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
	result := database.DB.Create(&products)
	if result.Error != nil{
		log.Fatalf("pfft what more can i say its not there: %v",  result.Error)
	}
	
	// fill the join table 
	database.DB.Exec("DELETE FROM customer_service")	
	for i := range customers{
		customer := &customers[i]

		// fill random services
		for i := 0; i<3; i++ {
			service := services[rand.Intn(len(services))]

			err := database.DB.Model(customer).Association("Services").Append(&service)
				
			if err != nil {
				log.Fatal(err)
				}
		}
		
	}
err := database.DB.AutoMigrate(&model.Customer{}, &model.Service{}, &model.Product{})
if err != nil{
	log.Fatal(err)
	}

}
