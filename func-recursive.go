package main

import "fmt"

func factorial(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * factorial(angka-1)
	}
}

func main() {
	fmt.Println(factorial(5))
}
