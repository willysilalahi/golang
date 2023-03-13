package main

import "fmt"

func main() {
	var name string

	name = "Eko Kurnia Marjan"
	fmt.Println(name)
	name = "Eko Kurnia Pohon Pinang"
	fmt.Println(name)

	var panggilan = "Samsoel"
	fmt.Println(panggilan)

	var age = 28
	fmt.Println(age)

	gelar := "Sarjana Teknik Informatika"
	sudah_tua := true
	fmt.Println(gelar)
	fmt.Println(sudah_tua)

	var (
		ayah = "Sutarjo"
		ibu  = "Suliswati"
	)
	fmt.Println("Ayah =", ayah)
	fmt.Println("Ibu =", ibu)
}
