package main

func main() {
	//cards := []string{NewCard(), "Ace of Diamond"}
	/*cards := deck{NewCard(), "Ace of Diamond"} /* here we created new var "cards" from type deck,
	since cards var has been created as type deck this var has an acces to function print that has been created in deck.go */

	//cards := newDeck() // here we call function which will create and retrung value car that has been created from two value from two slices

	//cards.saveToFile("my_cards") // Create a file with list of Cards

	cards := newDeckFromFile("my_cards")
	cards.print()

}
