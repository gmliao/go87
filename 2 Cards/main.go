package main

import "fmt"

func main() {

	// Create a new deck of cards
	cards := newDeck()

	// Shuffle the cards
	cards.shuffle()

	// Deal 5 cards to the hand
	hand_cards, cards := cards.deal(5)

	// Print the hand cards
	fmt.Println("Hand cards:")
	hand_cards.print()

	// Print the remaining cards
	fmt.Println("Remaining cards:", cards.length())

	// Save the remaining cards to a file
	cards.saveToFile("my_cards")

	// Load the cards from a file
	fmt.Println("Load cards:")
	load_cards := loadFromFile("my_cards")

	// Print the loaded cards
	load_cards.print()
}
