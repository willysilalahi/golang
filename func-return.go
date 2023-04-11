package main

import "fmt"
 

func getNilai(nilai int) string {
	if nilai >= 80 {
		return fmt.Sprint("Sangat Baik, Nilai anda = ", nilai)
	} else if nilai >= 60 && nilai < 80 {
		return fmt.Sprint("Baik, Nilai anda = ", nilai)
	} else if nilai >= 50 && nilai < 60 {
		return fmt.Sprint("Cukup, Nilai anda = ", nilai)
	} else {
		return fmt.Sprint("Buruk, Nilai anda = ", nilai)
	}
} 

func getTotal(price int, disc int) int{
	var harga int = price
	var diskon int = disc
	var total int
	total = harga - (harga * diskon/100)
	return total 
}

func main() { 
	fmt.Println(getNilai(89))
	fmt.Println("Total akhir dari harga = Rp.500.000 dan diskon = 15% adalah : ", getTotal(500000, 15))
}
