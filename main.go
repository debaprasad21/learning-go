package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

	abc := deck{"abc", "def"}
	abc.print()

}

func newCard() string {
	return "Five of Diamonds"
}
