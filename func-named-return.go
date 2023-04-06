package main

import "fmt" 

func getFullName()(firstName string, middleName string, lastName string){
	firstName = "Abdul"
	middleName = "Rachman"
	lastName = "Saragih"
	return
}

func main() { 
	awal, tengah, akhir := getFullName() 
	fmt.Println(awal)
	fmt.Println(tengah)
	fmt.Println(akhir)
}
