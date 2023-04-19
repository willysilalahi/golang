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
	hasil := Pembagian(10,2)
	fmt.Println(hasil)
}