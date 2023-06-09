package main
import "fmt"

type HasName interface{
	GetName() string
}

func sayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return "Anjingg"
}

func main(){
	person := Person{Name: "Tama"}
	sayHello(person)
}