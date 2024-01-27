package main

import "fmt"

type deck []string

func (d deck) print() {
	fmt.Println(d)
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, shape := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+shape)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
