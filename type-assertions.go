package main
import "fmt"

func random() interface{}{
	return "OK"
}

func  main()  {
	result := random()
	fmt.Println(result)
}