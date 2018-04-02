package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace of Diamonds", NewCard()}
	cards = append(cards, "Six of Shades")

	for i, cards := range cards {
		fmt.Println(i, cards)
	}
}

//NewCard is This func is test
func NewCard() string {
	return "Five of Diamonds"
}
