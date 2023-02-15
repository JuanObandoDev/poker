package test

import (
	"fmt"
	"testing"

	"github.com/JuanObandoDeveloper/poker/internal/app"
	"github.com/JuanObandoDeveloper/poker/internal/models"
)

func TestNewPoker(t *testing.T) {
	poker := app.NewPoker()
	if poker == nil {
		t.Errorf("Poker instance fails, got %v, want not nil", poker)
	}
}

func TestDistribute(t *testing.T) {
	poker := app.NewPoker()
	poker.Distribute()
	if len(*poker.PlayerDeck) != 5 {
		t.Errorf("Distribute fails, got %v, want 5", len(*poker.PlayerDeck))
	}
}

func TestDiscardDeck(t *testing.T) {
	poker := app.NewPoker()
	poker.Distribute()
	tables := []struct {
		indexesToDiscard []uint8
	}{
		{[]uint8{0, 1, 2, 3, 4}},
		{[]uint8{0, 1, 2, 3}},
		{[]uint8{0, 1, 2}},
		{[]uint8{0, 1}},
		{[]uint8{0}},
		{[]uint8{}},
	}
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		if len(*poker.PlayerDeck) != 5 {
			t.Errorf("DiscardDeck fails, got %v, want %v", len(*poker.PlayerDeck), 5)
		}
	}
}

func TestIsRoyalFlush(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Jack},
				{Suit: models.Hearts, Rank: models.Queen},
				{Suit: models.Hearts, Rank: models.King},
				{Suit: models.Hearts, Rank: models.Ace},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Diamonds, Rank: models.Jack},
				{Suit: models.Diamonds, Rank: models.Queen},
				{Suit: models.Diamonds, Rank: models.King},
				{Suit: models.Diamonds, Rank: models.Ace},
			},
		},
	}

	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		fmt.Println(poker.PlayerDeck)
		royalFlush := poker.IsRoyalFlush()
		if royalFlush != true {
			t.Errorf("IsRoyalFlush fails, got %v, want true", royalFlush)
		}
	}
}

func TestIsStraightFlush(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Hearts, Rank: models.Three},
				{Suit: models.Hearts, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Five},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Diamonds, Rank: models.Eight},
				{Suit: models.Diamonds, Rank: models.Nine},
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
	}

	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		straightFlush := poker.IsStraightFlush()
		if straightFlush != true {
			t.Errorf("IsStraightFlush fails, got %v, want true", straightFlush)
		}
	}
}

func TestIsPoker(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Two},
				{Suit: models.Spades, Rank: models.Two},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Seven},
				{Suit: models.Clubs, Rank: models.Seven},
				{Suit: models.Spades, Rank: models.Seven},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		poker := poker.IsPoker()
		if poker != true {
			t.Errorf("IsPoker fails, got %v, want true", poker)
		}
	}
}

func TestIsFullHouse(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Two},
				{Suit: models.Spades, Rank: models.Three},
				{Suit: models.Hearts, Rank: models.Three},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Seven},
				{Suit: models.Clubs, Rank: models.Seven},
				{Suit: models.Spades, Rank: models.Jack},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		fullHouse := poker.IsFullHouse()
		if fullHouse != true {
			t.Errorf("IsFullHouse fails, got %v, want true", fullHouse)
		}
	}
}

func TestIsFlush(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Hearts, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Six},
				{Suit: models.Hearts, Rank: models.Eight},
				{Suit: models.Hearts, Rank: models.Ten},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Ace},
				{Suit: models.Diamonds, Rank: models.Three},
				{Suit: models.Diamonds, Rank: models.Five},
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Diamonds, Rank: models.Nine},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		flush := poker.IsFlush()
		if flush != true {
			t.Errorf("IsFlush fails, got %v, want true", flush)
		}
	}
}

func TestIsStraight(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Three},
				{Suit: models.Clubs, Rank: models.Four},
				{Suit: models.Spades, Rank: models.Five},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Eight},
				{Suit: models.Clubs, Rank: models.Nine},
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		straight := poker.IsStraight()
		if straight != true {
			t.Errorf("IsStraight fails, got %v, want true", straight)
		}
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Two},
				{Suit: models.Spades, Rank: models.Five},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Seven},
				{Suit: models.Clubs, Rank: models.Seven},
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		threeOfAKind := poker.IsThreeOfAKind()
		if threeOfAKind != true {
			t.Errorf("IsThreeOfAKind fails, got %v, want true", threeOfAKind)
		}
	}
}

func TestIsTwoPair(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Four},
				{Suit: models.Spades, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Hearts, Rank: models.Seven},
				{Suit: models.Clubs, Rank: models.Nine},
				{Suit: models.Spades, Rank: models.Nine},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		twoPair := poker.IsTwoPair()
		if twoPair != true {
			t.Errorf("IsTwoPair fails, got %v, want true", twoPair)
		}
	}
}

func TestIsOnePair(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Ace},
				{Suit: models.Diamonds, Rank: models.Ace},
				{Suit: models.Clubs, Rank: models.Four},
				{Suit: models.Spades, Rank: models.Five},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			[]models.Card{
				{Suit: models.Diamonds, Rank: models.Jack},
				{Suit: models.Hearts, Rank: models.Jack},
				{Suit: models.Clubs, Rank: models.Nine},
				{Suit: models.Spades, Rank: models.Ten},
				{Suit: models.Diamonds, Rank: models.Seven},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		onePair := poker.IsOnePair()
		if onePair != true {
			t.Errorf("IsOnePair fails, got %v, want true", onePair)
		}
	}
}

func TestValidate(t *testing.T) {
	poker := app.NewPoker()
	tables := []struct {
		pokerDeck models.Deck
	}{
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Ten},
				{Suit: models.Hearts, Rank: models.Jack},
				{Suit: models.Hearts, Rank: models.Queen},
				{Suit: models.Hearts, Rank: models.King},
				{Suit: models.Hearts, Rank: models.Ace},
			},
		},
		{
			models.Deck{
				{Suit: models.Diamonds, Rank: models.Seven},
				{Suit: models.Diamonds, Rank: models.Eight},
				{Suit: models.Diamonds, Rank: models.Nine},
				{Suit: models.Diamonds, Rank: models.Ten},
				{Suit: models.Diamonds, Rank: models.Jack},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Two},
				{Suit: models.Spades, Rank: models.Two},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Two},
				{Suit: models.Spades, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Four},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Hearts, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Six},
				{Suit: models.Hearts, Rank: models.Eight},
				{Suit: models.Hearts, Rank: models.Ten},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Three},
				{Suit: models.Clubs, Rank: models.Four},
				{Suit: models.Spades, Rank: models.Five},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Two},
				{Suit: models.Spades, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Two},
				{Suit: models.Clubs, Rank: models.Four},
				{Suit: models.Spades, Rank: models.Four},
				{Suit: models.Hearts, Rank: models.Six},
			},
		},
		{
			models.Deck{
				{Suit: models.Hearts, Rank: models.Two},
				{Suit: models.Diamonds, Rank: models.Four},
				{Suit: models.Clubs, Rank: models.Six},
				{Suit: models.Spades, Rank: models.Jack},
				{Suit: models.Hearts, Rank: models.Jack},
			},
		},
	}
	for _, table := range tables {
		poker.PlayerDeck = &table.pokerDeck
		win := poker.Validate()
		if win != true {
			t.Errorf("Validate fails, got %v, want true", win)
		}
	}
}
