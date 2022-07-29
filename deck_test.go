package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != len(cardSuits)*len(cardValues) {
		t.Errorf(
			"Expected deck length of %d, but got %d",
			len(cardSuits)*len(cardValues),
			len(cards),
		)
	}

	if cards[0] != cardValues[0]+cardSuits[0] {
		t.Errorf(
			"Expected first card to be %v, but got %v",
			cardValues[0]+cardSuits[0],
			cards[0],
		)
	}

	if cards[len(cards)-1] != cardValues[len(cardValues)-1]+cardSuits[len(cardSuits)-1] {
		t.Errorf(
			"Expected last card to be %v, but got %v",
			cardValues[len(cardValues)-1]+cardSuits[len(cardSuits)-1],
			cards[len(cards)-1],
		)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("files/_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(cardSuits)*len(cardValues) {
		t.Errorf("Error!!!")
	}

	os.Remove("files/_decktesting")
}
