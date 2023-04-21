package main
import (
	"fmt"
	"errors"	
)
	

func Pembagian(value int, pembagi int) (int, error){
	if pembagi == 0{
		return 0, errors.New("Bilangan ndak bisa di bagi 0 bangsat!")
	}else{
		return value/pembagi, nil
	}
}


func main(){
	
	var controller error = errors.New("Error ya!")
	fmt.Println(controller.Error())

	hasil, err := Pembagian(10,0)
	if err == nil{
		fmt.Println("Hasil =", hasil)
	}else{
		fmt.Println("Error :", err.Error())
	}
}