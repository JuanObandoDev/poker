package test

import (
	"testing"

	"github.com/JuanObandoDeveloper/poker/internal/models"
)

func TestNewMainDeck(t *testing.T) {
	deck := models.NewMainDeck()
	if len(*deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(*deck))
	}
}

func TestShuffle(t *testing.T) {
	deck := models.NewMainDeck()
	firstCards := deck.Deck(0, 1, 2, 3, 4)
	deck.Shuffle()
	firstShuffledCards := deck.Deck(0, 1, 2, 3, 4)
	for i := 0; i < 5; i++ {
		if (*firstCards)[i] == (*firstShuffledCards)[i] {
			t.Errorf("Expected %v to be different from %v", (*firstCards)[i], (*firstShuffledCards)[i])
		}
	}
}

func TestDealDeck(t *testing.T) {
	deck := models.NewMainDeck()
	deck.Shuffle()
	deckToDeal := deck.DealDeck(5)
	if len(*deckToDeal) != 5 {
		t.Errorf("Expected deck length of 5, but got %v", len(*deckToDeal))
	}
}

func TestAddDeck(t *testing.T) {
	deck := models.NewMainDeck()
	deck.Shuffle()
	deckToDeal := deck.DealDeck(5)
	deck.AddDeck(deckToDeal)
	if len(*deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(*deck))
	}
}

func TestRemoveDeck(t *testing.T) {
	deck := models.NewMainDeck()
	deck.Shuffle()
	deckToDeal := deck.DealDeck(5)
	deck.RemoveDeck(deckToDeal)
	if len(*deck) != 47 {
		t.Errorf("Expected deck length of 47, but got %v", len(*deck))
	}
}

func TestDeck(t *testing.T) {
	deck := models.NewMainDeck()
	deck.Shuffle()
	deckToDeal := deck.Deck(0, 1, 2, 3, 4)
	if len(*deckToDeal) != 5 {
		t.Errorf("Expected deck length of 5, but got %v", len(*deckToDeal))
	}
}

func TestSort(t *testing.T) {
	deck := models.NewMainDeck()
	deck.Shuffle()
	firstShuffledCards := deck.Deck(0, 1, 2, 3, 4)
	unsortedCards := firstShuffledCards.Deck(0, 1, 2, 3, 4)
	firstShuffledCards.Sort()
	for i := 0; i < 5; i++ {
		if (*firstShuffledCards)[i] == (*unsortedCards)[i] {
			t.Errorf("Expected %v to be different to %v", (*firstShuffledCards)[i], (*unsortedCards)[i])
		}
	}
}
