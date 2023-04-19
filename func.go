package main

import "fmt"

func sayHello() {
	fmt.Println("Hello world")
}

func getStatus(nilai int) {
	if nilai >= 80 {
		fmt.Println("Sangat Baik")
	} else if nilai >= 60 && nilai < 80 {
		fmt.Println("Baik")
	} else if nilai >= 50 && nilai < 60 {
		fmt.Println("Cukup")
	} else {
		fmt.Println("Buruk")
	}
}

func main() {
	var hsl float32

	hsl = 0.1
	fmt.Println(hsl)
}
