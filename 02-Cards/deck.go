package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 建立一個新的型別，叫做 deck，它是一個 string 的 slice
type deck []string

// print prints out all the cards in the deck
func (d deck) print() {
	// Iterate through every card in the deck
	for index, card := range d {
		// Print the card number and card
		fmt.Println(index, card)
	}
}

func (d deck) length() int {
	return len(d)
}

// deal deals a hand of cards
func (d deck) deal(handSize int) (deck, deck) {
	// Return the first handSize cards and the rest of the cards
	return d[:handSize], d[handSize:]
}

// toString converts the deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile saves the deck to a file
func (d deck) saveToFile(filename string) error {
	serializedDeck := []byte(d.toString())
	return os.WriteFile(filename, []byte(serializedDeck), 0666)
}

// loadFromFile loads a deck from a file
func loadFromFile(filename string) deck {
	// Read the file
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return deck{}
	}

	// Convert the file to a string
	return deck(strings.Split(string(bs), ","))
}

// shuffle shuffles the deck
func (d deck) shuffle() {
	// Create a source for random numbers
	randSource := rand.NewSource(time.Now().UnixNano())
	// Create a random number generator from the source
	r := rand.New(randSource)
	// Iterate through the deck
	for index := range d {
		// Generate a random number between 0 and the length of the deck minus 1
		newPosition := r.Intn(len(d) - 1)
		// Swap the current card with the card at the random position
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

// newDeck creates a new deck of cards
func newDeck() deck {
	// Create a new deck
	cards := deck{}

	// Create a slice of strings for the suits
	cardSuits := []string{"♠", "♦", "♥", "♣"}

	// Create a slice of strings for the values
	cardValues := []string{"1", "2", "3", "4", "5"}

	// Iterate through the suits
	for _, suit := range cardSuits {
		// Iterate through the values
		for _, value := range cardValues {
			// Append the value and suit to the deck
			cards = append(cards, suit+value)
		}
	}

	// Return the deck
	return cards
}
