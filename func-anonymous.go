package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist1 := func(name string) bool {
		return name == "admin"
	}

	blacklist2 := func(name string) bool {
		return name == "adi"
	}

	registerUser("antony", blacklist1)
	registerUser("edwin", blacklist2)
}
