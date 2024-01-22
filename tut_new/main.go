package main

import "fmt"

func main() {
	cards := []string{"Card1", "Card2", "Card3"}
	cards = append(cards, "Card4")
	cards = append(cards, newCard())
	for ind, card := range cards {
		fmt.Println(ind, card)
	}
}
func newCard() string {
	return "A Plain Card"
}
