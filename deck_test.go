package main

import (
	"os"
	"testing"
)

// TestNewDeck is a function that tests the newDeck function
// it takes a testing.T object as an argument
// the testing.T object is used to report and log errors
// the function name must start with Test
// go test will run this function
// Errorf is a function on the testing.T object
func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Check if the length of the deck is 16
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Check if the first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Check if the last card is Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingCards := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand length of 5, but got %v", len(hand))
	}

	if len(remainingCards) != 11 {
		t.Errorf("Expected remaining cards length of 11, but got %v", len(remainingCards))
	}
}

// TestSaveToFileAndNewDeckFromFile is a function that tests the saveToFile and newDeckFromFile functions
// os.Remove("_decktesting") -> remove the file before testing
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Remove the file before testing
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
