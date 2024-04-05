package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
	fmt.Println("----------")
	fmt.Println(newDeckFromFile("my_cards"))
}
