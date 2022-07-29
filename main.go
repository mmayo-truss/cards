package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards)

	cards.shuffle()

	fmt.Println(cards)
}
