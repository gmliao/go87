package main

import (
	"os"
	"testing"
)

var expectedLength int = 5 * 4
var firstCard string = "♠1"
var lastCard string = "♣5"

func TestNewDeck(t *testing.T) {
	// Create a new deck of cards
	cards := newDeck()

	// Check the number of cards
	if len(cards) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(cards))
	}

	// Check the first card
	if cards[0] != firstCard {
		t.Errorf("Expected first card of Ace of Spades, but got %v", cards[0])
	}

	// Check the last card
	if cards[len(cards)-1] != lastCard {
		t.Errorf("Expected last card of King of Clubs, but got %v", cards[len(cards)-1])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	var testFile string = "_decktesting"

	// Create a new deck of cards
	cards := newDeck()

	// Save the cards to a file
	e := cards.saveToFile(testFile)
	if e != nil {
		t.Errorf("Error saving file %v", testFile)
	}

	// Load the cards from a file
	loadedCards := loadFromFile(testFile)

	// Check the number of cards
	if len(loadedCards) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(loadedCards))
	}

	// Check the first card
	if loadedCards[0] != firstCard {
		t.Errorf("Expected first card of Ace of Spades, but got %v", loadedCards[0])
	}

	// Check the last card
	if loadedCards[len(loadedCards)-1] != lastCard {
		t.Errorf("Expected last card of King of Clubs, but got %v", loadedCards[len(loadedCards)-1])
	}

	// Remove the test file if it exists
	e = os.Remove(testFile)
	if e != nil {
		t.Errorf("Error removing file %v", testFile)
	}
}
