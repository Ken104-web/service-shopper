package cli

import (
	"fmt"
	"log"
	"service-shopper/database"
	"service-shopper/model"

	"github.com/spf13/cobra"
)

var shopperCmd = &cobra.Command{
	Use: "Ujenzi Market",
	Short: "A simple display CLI",
	Long: "Ujenzi Market is a CLI app use to display various aspect of the establishmet such as num of customers services and products",

	// display customers
	Run: func(cmd *cobra.Command, args []string) {
		database.Data()

		fmt.Println("Here is alist of all customers: ")

		var customers []model.Customer
		r := database.DB.Find(&customers)
		if r.Error != nil{
			fmt.Println("Err fetching customers: ", r.Error)
		}
		for i, c := range customers{
			fmt.Printf("%d. %s\n", i+1, c.Name)
		}

		// display through query
		var name string
		fmt.Println("\n Enter customer name: ")
		fmt.Scanln(&name)

		var c model.Customer
		r = database.DB.Where("name =  ?", name).First(&c)

		if r.Error != nil{
			fmt.Println("Customer not found")
			return
		}
		fmt.Printf("\nCustomer Name: %s\n", c.Name)

	},


}

func Execute(){
	if err := shopperCmd.Execute(); err != nil{
		log.Fatal(err)
	}
}
