package main

import "fmt"
 
func getSum(name string, number ...int) (string, int) {
	total := 0

	for _, value := range number {
		total+= value
	}
	total = total/len(number)
	return name, total
} 

func main() {
	name, nilai := getSum("Edwin", {90, 89, 88, 95, 82})
	fmt.Println("Mendapatkan rata-rata nilai : ", nilai)
	fmt.Println("Mendapatkan rata-rata nilai : ", name)
}
