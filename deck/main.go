package main

func main() {
	cards := newDeckFromFile("go_decklist")
	//cards := newDeck()
	//cards.saveToFile("go_decklist")
	cards.shuffle()
	cards.print()
}
