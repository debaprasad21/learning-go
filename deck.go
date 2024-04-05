package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// We are not using the index value so we can replace it with _
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	// slice syntax
	// d[:handSize] -> get a slice from the start of the slice to the handSize (not including the handSize)
	// d[handSize:] -> get a slice from the handSize to the end of the slice
	return d[:handSize], d[handSize:]
}

// newDeckFromFile is a function that returns a deck
// it takes a string as an argument the filename
// it reads the file and returns a deck
// the deck is created by splitting the string by the comma post reading the file converting bytes to string
func newDeckFromFile(filename string) deck {
	// bs is a slice of bytes
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		// os.Exit(1) -> exit the program and 1 is the status code that represents an error
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// the print function is a receiver function
// it is a function that is attached to a type
// in this case, the type is deck
// the receiver function is called on a variable of type deck
// the variable is available in the function as a variable called 'd'
// the receiver function is called like this: cards.print()
// it refers to this in js or otherwise we can call it as an instance in reference to Object Oriented Programming
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// toString is a receiver function
// it is a function that is attached to a type
// converting deck -> []string -> string
func (d deck) toString() string {
	// type conversion
	return strings.Join([]string(d), ",")
}

// the error type is a built-in type in go - it is a type that represents an error
// 0666 is a permission level that represents read and write access
// saveToFile is a receiver function
// saving the cards to a file post converting deck -> []string -> string -> []byte
func (d deck) saveToFile(filename string) {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// shuffle is a receiver function
// uses the math/rand package to shuffle the deck
// it generates a random number and swaps the values
// the rand.Intn() function generates a random number
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// generating a random number
		newPosition := r.Intn(len(d) - 1)
		// swapping the values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
