package main
import "fmt"

type Customer struct{
	Name, Address string
	Age int
	Phone string
}

func  main()  {
	var cust1 Customer
	// Assign 1 
	cust1.Name = "Suliswati"
	cust1.Address = "Percut"
	cust1.Age = 39
	fmt.Println(cust1)
	fmt.Println(cust1.Name)
	fmt.Println(cust1.Age)

	// Assign 2
	 cust2 := Customer{
		Name: "Agnes",
		Address: "Marelan",
		Age: 23,
	}
	fmt.Println(cust2)
	
	// Assign 3 
	var cust3 = Customer{"Hari Tod", "Tj. Morawa", 25, "082188781181"}
	fmt.Println(cust3)
}