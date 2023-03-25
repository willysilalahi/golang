package main

import "fmt"

func main() {

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// var slice1 = months[4:7]

	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// slice1[0] = "Neptunus"

	// fmt.Println(slice1)
	fmt.Println(months)
	// fmt.Println("================")

	// slice2 := months[8:10]
	// slice2 = append(slice2, "Pluto")
	// slice2 = append(slice2, "Merkurius")
	// slice2 = append(slice2, "SimonPetrus")
	// fmt.Println(months)
	// fmt.Println(slice2)

	// slice3 := months[10:]
	// slice3[0] = "Pluto"
	// slice4 := append(slice3, "Merkurius")
	// slice4[0] = "SimonPetrus"
	// fmt.Println(months)
	// fmt.Println(slice3)
	// fmt.Println(slice4)

	newSlice := make([]string, 2, 8)
	newSlice[0] = "Suliswati"
	newSlice[1] = "Sembiring"
	newSlice = append(newSlice, "S.Pd")
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{2, 4, 6, 8, 10}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
