package main

import "fmt"

func getGoodBye(name string) string{
	return "Good Bye " + name + " Anjenggg"
}


func  main()  {
	sayHello := getGoodBye

	fmt.Println(sayHello("Rustam"))
}