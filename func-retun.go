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

func main() {
	sayHello()
	getStatus(45)
	fmt.Println(getNilai(56)) 
}
