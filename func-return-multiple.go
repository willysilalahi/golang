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

func getTotalPrice(price int, disc int) (int, int, string){
	diskon := price * disc / 100
	concatenated := fmt.Sprint(disc, "%")
	return price, diskon, concatenated
}

func main() {  
	price, totalDisc, disc := getTotalPrice(500000,15)
	fmt.Println("Total diskon dari harga", price, "dan diskon", disc, "adalah :", totalDisc) 
}
