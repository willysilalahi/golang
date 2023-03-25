package main

import "fmt"

func main() {

	var angka int32 = 190
	var bilangan int64 = int64(angka)

	fmt.Println(angka)
	fmt.Println(bilangan)

	nama := "Riyan Kurniawan"
	var tmp = nama[0]
	var newtmp = string(tmp)

	fmt.Println(nama)
	fmt.Println(tmp)
	fmt.Println(newtmp)
}
