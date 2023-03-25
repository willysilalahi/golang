package main

import "fmt"

func  main()  {
	
	// for i := 0; i < 20; i++ {
	// 	if(i < 10){
	// 		for j := 10-i ; j > 0; j-- {
	// 			fmt.Print("*")
	// 		}	
	// 	}else{ 			
	// 		for j := 0 ; j < i - (10 -1); j++ {
	// 			fmt.Print("*")
	// 		} 
	// 	}
	// 	fmt.Println("")
	// }



	// counter := 1;
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }



	slice1 := []string{"Steven", "Hari", "Yudi", "Edwin", "Riyan", "Agus"}

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	fmt.Println("-------------------")

	for i, _ := range slice1{
		//fmt.Println("Index", i, " =", value)
		fmt.Println("Index =", i)
	}

	
}