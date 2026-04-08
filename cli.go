package main

import (
	"fmt"
	"service-shopper/model"

	"github.com/spf13/cobra"
)

var shopperCmd = &cobra.Command{
	Use: "Ujenzi Market",
	Short: "A simple display CLI",
	Long: "Ujenzi Market is a CLI app use to display various aspect of the establishmet such as num of customers services and products",

	// display customers
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here is alist of all customers: ")

		// var customers []model.Customer
	},
	


}
