package main

import "fmt"

func main() {
	var person = "Farkhod"
	fmt.Println(printhello(), person)
	fmt.Println(number())
}

func printhello() string {
	return "Hello :"
}

func number() int {
	return 123123
}
