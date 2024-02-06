package main

func main() {
	cards := newDeck()
	// hand, remaining := deal(cards, 5)
	cards.saveToFile("my_cards")
	hand, remaining := deal(cards, 5)
}
