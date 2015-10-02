package deck

import "math/rand"

//Deck is the struct in which card structs are stored and deck methods are defined
type Deck struct {
	cards []Card
}

//Card is the struct representation of a playing card, with a suite, name and value
type Card struct {
	Suite string
	Name  string
	Value int
}

//DealCard returns a single Card to the user and removes that card from the deck
func (deck *Deck) DealCard() Card {
	randomCardNumber := rand.Intn(deck.Size())
	randomCard := deck.cards[randomCardNumber]
	deck.cards = append(deck.cards[:randomCardNumber], deck.cards[randomCardNumber+1:]...)
	return randomCard
}

//Size returns the amount of cards left in the deck
func (deck *Deck) Size() int {
	return len(deck.cards)
}

//New returns a pointer to a new 52 card standard playing deck
func New() (deckPointer *Deck) {
	deck := Deck{}

	suites := [4]string{"Spades", "Clubs", "Hearts", "Diamonds"}
	for i := range suites {
		names := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
		values := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}
		for j := range values {
			deck.cards = append(deck.cards, Card{suites[i], names[j], values[j]})
		}
	}

	return &deck
}
