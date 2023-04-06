package main

import "fmt"
 
func getResult(nilai int) (string, int) {
	if nilai >= 80 {
		return "Sangat Baik, Nilai anda =", nilai
	} else if nilai >= 60 && nilai < 80 {
		return "Baik, Nilai anda =", nilai
	} else if nilai >= 50 && nilai < 60 {
		return "Cukup, Nilai anda =", nilai
	} else {
		return "Buruk, Nilai anda =", nilai
	}

} 

func main() { 
	teks, nilai := getResult(98)
	fmt.Println(teks, nilai) 
}
