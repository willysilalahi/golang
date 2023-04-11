package main

import "fmt" 

func getFullName()(firstName string, middleName string, lastName string){
	firstName = "Bertin"
	middleName = "Rachman"
	lastName = "Saragih"
	return
}

func getNilaiAkhir(tugas int, uts int, uas int)(resutlTugas int, resultUTS int, resultUAS int, total int) {
	// Tugas = 20%, UTS = 35, UAS = 45
	resutlTugas = tugas * 20 / 100
	resultUTS = uts * 35 / 100
	resultUAS = uas * 45 / 100
	total = resutlTugas + resultUTS + resultUAS
	return
}

func main() { 
	awal, tengah, akhir := getFullName() 
	fmt.Println(awal)
	fmt.Println(tengah)
	fmt.Println(akhir)
	fmt.Println("-----------------")
	tugas, uts, uas, total := getNilaiAkhir(89,85,95) 
	fmt.Println("Tugas =", tugas)
	fmt.Println("UTS =", uts)
	fmt.Println("UAS =", uas)
	fmt.Println("Nilai Akhir =", total)
}
