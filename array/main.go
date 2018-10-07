package main

import (
	"fmt"
)

func main() {
	card := []string{"This is first line", newCard()}
	card = append(card, "This is first second")
	fmt.Println(card)
	for i, data := range card {
		fmt.Println(i, data)
	}
}

func newCard() string {
	return "This is first third"
}
