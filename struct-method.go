package main
import "fmt"

type Customer struct{
	Name, Address string
	Age int
	Phone string
}

//func biasa
func sayHello(customer Customer, name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}

//struct method
func (customer Customer) sayHi(name string){
	fmt.Println("Hii", name, "my name is", customer.Name)
}

func  main()  {	
	 cust1 := Customer{
		Name: "Agnes",
		Address: "Marelan",
		Age: 23,
	} 
	sayHello(cust1, "Ruth")

	cust2 := Customer{
		Name: "Steven",
		Address: "Brayan",
		Age: 23,
	} 
	cust2.sayHi("Yolanda")
}