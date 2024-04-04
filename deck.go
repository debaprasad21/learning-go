package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// the print function is a receiver function
// it is a function that is attached to a type
// in this case, the type is deck
// the receiver function is called on a variable of type deck
// the variable is available in the function as a variable called 'd'
// the receiver function is called like this: cards.print()
// it refers to this or otherwise we can call it as an instance inreference to Object Oriented Programming
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card, d)
	}
}
