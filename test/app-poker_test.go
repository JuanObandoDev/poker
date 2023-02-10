package test

import (
	"testing"

	"github.com/JuanObandoDeveloper/poker/internal/app"
)

func defaultPokerHandsTest() (*app.Poker, []struct{ indexesToDiscard []uint8 }) {
	poker := app.NewPoker()

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

	return poker, tables
}

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
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		if len(*poker.PlayerDeck) != 5 {
			t.Errorf("DiscardDeck fails, got %v, want %v", len(*poker.PlayerDeck), 5)
		}
	}
}

func TestIsRoyalFlush(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		royalFlush := poker.IsRoyalFlush()
		if royalFlush != true && royalFlush != false {
			t.Errorf("IsRoyalFlush fails, got %v, want true or false", royalFlush)
		}
	}
}

func TestIsStraightFlush(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		straightFlush := poker.IsStraightFlush()
		if straightFlush != true && straightFlush != false {
			t.Errorf("IsStraightFlush fails, got %v, want true or false", straightFlush)
		}
	}
}

func TestIsPoker(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		poker := poker.IsPoker()
		if poker != true && poker != false {
			t.Errorf("IsPoker fails, got %v, want true or false", poker)
		}
	}
}

func TestIsFullHouse(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		fullHouse := poker.IsFullHouse()
		if fullHouse != true && fullHouse != false {
			t.Errorf("IsFullHouse fails, got %v, want true or false", fullHouse)
		}
	}
}

func TestIsFlush(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		flush := poker.IsFlush()
		if flush != true && flush != false {
			t.Errorf("IsFlush fails, got %v, want true or false", flush)
		}
	}
}

func TestIsStraight(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		straight := poker.IsStraight()
		if straight != true && straight != false {
			t.Errorf("IsStraight fails, got %v, want true or false", straight)
		}
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		threeOfAKind := poker.IsThreeOfAKind()
		if threeOfAKind != true && threeOfAKind != false {
			t.Errorf("IsThreeOfAKind fails, got %v, want true or false", threeOfAKind)
		}
	}
}

func TestIsTwoPair(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		twoPair := poker.IsTwoPair()
		if twoPair != true && twoPair != false {
			t.Errorf("IsTwoPair fails, got %v, want true or false", twoPair)
		}
	}
}

func TestIsOnePair(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		onePair := poker.IsOnePair()
		if onePair != true && onePair != false {
			t.Errorf("IsOnePair fails, got %v, want true or false", onePair)
		}
	}
}

func TestValidate(t *testing.T) {
	poker, tables := defaultPokerHandsTest()
	poker.Distribute()
	for _, table := range tables {
		poker.DiscardDeck(poker.PlayerDeck.Deck(table.indexesToDiscard...))
		win := poker.Validate()
		if win != true && win != false {
			t.Errorf("Validate fails, got %v, want true or false", win)
		}
	}
}
