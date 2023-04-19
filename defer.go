package main

import "fmt"

func logging() {
	fmt.Println("Logging di panggil ya!")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Application Running!")
	fmt.Println("Updated a package!")
	fmt.Println("Updated a package!")
	fmt.Println("Updated a package!")
	fmt.Println("Updated a package!")
	fmt.Println("Updated a package!")
	fmt.Println("Updated a package!")
	fmt.Println("Running packages!")
	result := 10 / value
	fmt.Println("Result =", result)
}

func main() {
	runApp(10)
}
