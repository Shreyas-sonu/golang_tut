package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}
func newCard() string {
	return "A Plain Card"
}

// func newCard() int {
// 	return 1
// }
