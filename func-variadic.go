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
	hasil := []int{90, 89, 88, 95, 82}
	name, nilai := getSum("Edwin", hasil...)
	fmt.Println(name, "mendapatkan rata-rata nilai :", nilai) 
}
