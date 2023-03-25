package main

import "fmt"

func main() {

	var students [5]string
	students[0] = "Hari"
	students[1] = "Roby"
	students[2] = "Agung"
	students[3] = "Ferdi"
	students[4] = "Steven"

	fmt.Println(students)

	var nilai = [5]int{
		70,
		89,
		82,
		80,
		89,
	}
	nilai[4] = 92
	fmt.Println(nilai)
	fmt.Println(len(nilai))

	var data [10]string
	fmt.Println(len(data))

	var titik = [...]int{
		3,
		4,
		2,
		5,
		9,
		1,
	}
	fmt.Println(titik)
}
