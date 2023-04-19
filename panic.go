package main

import (
	"fmt"
)

func endApp() {
	msg := recover()
	if msg != nil {
		fmt.Println("Message nya", msg)
	}
	fmt.Println("Aplikasi Selesai!")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APPLICATION ERRORR!!")
	}
	fmt.Println("Aplikasi Berjalan")

}

func main() {
	runApp(false)
	//main point : Jika panic or function error, app tidak berhenti, tapi kalo revocer() di comment, println("Eko") tidak dapat dijalankan
	fmt.Println("Eko")
}
