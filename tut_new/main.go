package main

func main() {
	cards := newDeck()
	hand, remaining := deal(cards, 5)
	hand.print()
	remaining.print()
}
func newCard() string {
	return "Five of Diamonds"
}
