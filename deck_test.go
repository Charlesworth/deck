package deck

import (
	"reflect"
	"testing"
)

var suites = []string{"Spades", "Clubs", "Hearts", "Diamonds"}
var names = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

func TestNewReturnsDeck(t *testing.T) {
	if reflect.TypeOf(New()) != reflect.TypeOf(&Deck{}) {
		t.Error("NewDeck() does not return type *Deck")
	}
}

func TestNewContainsFiftyTwoCards(t *testing.T) {
	newDeck := New()
	newDeckSize := newDeck.Size()
	if newDeckSize != 52 {
		t.Error("NewDeck does not return a deck with 52 Cards, it contained", newDeckSize)
	}
}

func TestDeckSize(t *testing.T) {
	newDeck := New()
	newDeckSize := newDeck.Size()
	if newDeckSize != 52 {
		t.Error("Deck.Size() is reporting the incorrect deck size back, should be 52 for a new deck and is", newDeckSize)
	}

	oneCardDeck := Deck{cards: []Card{{"test", "test", 1}}}
	oneCardDeckSize := oneCardDeck.Size()
	if oneCardDeckSize != 1 {
		t.Error("Deck.Size() is returning the incorrect deck size back, should be 1 for a deck with 1 card and is", oneCardDeckSize)
	}
}

func TestDealCardReturnsTypeCard(t *testing.T) {
	testCard := New().DealCard()
	if reflect.TypeOf(testCard) != reflect.TypeOf(Card{}) {
		t.Error("Deck.DealCard did not return type Card")
	}
}

func TestDealCardReturnsACardWithValue(t *testing.T) {
	testCard := New().DealCard()
	if !stringSliceContains(testCard.Suite, suites) {
		t.Error("Deck.DealCard did not return a valid card suite:", testCard.Suite)
	}

	if !stringSliceContains(testCard.Name, names) {
		t.Error("Deck.DealCard did not return a valid card name:", testCard.Name)
	}

	if !(testCard.Value >= 1) || !(testCard.Value <= 10) {
		t.Error("Deck.DealCard did not return a valid card value:", testCard.Value)
	}
}

func stringSliceContains(element string, slice []string) bool {
	for i := range slice {
		if slice[i] == element {
			return true
		}
	}
	return false
}

func TestDealCardRemovesACardFromDeck(t *testing.T) {
	testDeck := New()
	testDeck.DealCard()
	deckWithOneCardDeltSize := testDeck.Size()
	if deckWithOneCardDeltSize != 51 {
		t.Error("Deck.DealCard() does not remove the card from the deck, size of deck should be 51, instead its", deckWithOneCardDeltSize)
	}
}

func TestDealCardReturnsAndRemovesSameCard(t *testing.T) {
	testDeck := New()
	removedCard := testDeck.DealCard()
	if deckContains(testDeck, removedCard) {
		t.Error("Deck.DealCard() does not remove the same card it returns to the funcation caller")
	}
}

func deckContains(deck *Deck, card Card) bool {
	for i := range deck.cards {
		if deck.cards[i] == card {
			return true
		}
	}
	return false
}
