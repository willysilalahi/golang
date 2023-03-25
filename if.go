package main

import "fmt"

func main() {
	nama := "Susi Susanti Julianti Sinaga"

	if leng := len(nama); leng > 10 {
		fmt.Println("Nama anda terlalu panjang, babi!")
	} else {
		fmt.Println("Udah mantep!")
	}
}
