package main

import (
	"fmt"
)

func main() {
	dataMap := map[string]string{
		"nama":          "Budi Ismail",
		"jenis_kelamin": "Laki-laki",
		"umur":          "12 Tahun",
	}

	fmt.Println(dataMap)
	dataMap["pekerjaan"] = "Programmer"

	newMap := make(map[string]string)
	fmt.Println(newMap)
	newMap["name"] = "Mouse"
	newMap["qty"] = "2"
	newMap["status"] = "Bagus"
	fmt.Println(newMap)
	delete(newMap, "status")
	fmt.Println(newMap)
	fmt.Println(len(newMap))
}
