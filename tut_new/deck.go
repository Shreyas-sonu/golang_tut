package main

import "fmt"

type deck []string

func (d deck) print() {
	fmt.Println(d)
	for i, card := range d {
		fmt.Println(i, card)
	}
}
