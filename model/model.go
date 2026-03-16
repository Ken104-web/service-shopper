package main

import (
	"gorm.io/gorm"
)
type Customer struct{
	gorm.Model
	Name string

	// one-to-many with Product
	Products []Product`gorm:"foreignKey:CustomerID;"`
}

type Service struct{
	gorm.Model
	Name string
}

type Product struct{
	gorm.Model
	Item_Name string
	
	CustomerID uint // foreignKey to customer 
}
