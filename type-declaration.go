package main

import "fmt"

func main() {
	type ktp string
	type nikah bool

	var noKtp ktp = "1208277192999001"
	var hasMarried nikah = true

	fmt.Println("No KTP =", noKtp)
	fmt.Println("Sudah menikah =", hasMarried)
}
