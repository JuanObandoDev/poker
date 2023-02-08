package models

type Suit int
type Rank int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (s Suit) String() string {
	return [...]string{"♠", "♥", "♦", "♣"}[s]
}

func (r Rank) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}[r]
}

var Suits = [...]Suit{Spades, Hearts, Diamonds, Clubs}
var Ranks = [...]Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return c.Rank.String() + c.Suit.String()
}
