package main

import "fmt"

func newMap(name string) map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	var data map[string]string = nil
	person := newMap("Natu")
	fmt.Println(data)
	fmt.Println(person)
}