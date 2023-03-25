package main

import "fmt"

func main(){
	status := 2

	switch status {
	case 0:
		fmt.Println("Order Created")
	case 1:
		fmt.Println("Order Confirmed")
	case 2:
		fmt.Println("Order Item On Packaging Proccess")
	case 3:
		fmt.Println("Order Item Ready for Shipping")
	case 4:
		fmt.Println("Order Item Sent to Customer")
	case 5:
		fmt.Println("Item Has Been arrived")
	default:
		fmt.Println("Order Done")
	}
	
}