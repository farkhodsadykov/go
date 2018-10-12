package main

import (
	"fmt"
)

func main() {
	card := []string{"This is first line", newCard()}
	card = append(card, "This is second line")
	// fmt.Println(card)
	for i, data := range card {
		fmt.Println(i, data)
	}
}

func newCard() string {
	return "This is third line"
}
