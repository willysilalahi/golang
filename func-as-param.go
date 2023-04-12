package main

import "fmt"

type Filter func(string) string

func getFilter(word string) string{
	if word == "Anjing" {
		return "..."
	}else if word == "Babi" {
		return "..."
	}else if word == "Monyet" {
		return "..."
	}else if word == "Kampret" {
		return "..."
	}else{
		return word
	}
}

func filterWords(name string, filter Filter){
	fmt.Println("Hello", filter(name))
}


func  main()  {
	filterWords("Rustam", getFilter)
}