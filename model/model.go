package model

import (
	"gorm.io/gorm"
)
type Customer struct{
	gorm.Model
	Name string

	// one-to-many with Product
	Products []Product`gorm:"foreignKey:CustomerID;"`

	// many-to-many with service 
	Services []Service`gorm:"many2many:customer_service;"`
}

type Service struct{
	gorm.Model
	Name string
	Rating int

	// one-to-many with Service
	Products []Product`gorm:"foreignKey:ServiceID;"`

	// many-to-many with customer 
	Customers []Customer`gorm:"many2many:customer_service;"`
}

type Product struct{
	gorm.Model
	Item_Name string
	Price int
	
	ServiceID uint // foreignKey to service 
	CustomerID uint // foreignKey to customer 
}
