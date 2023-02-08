package app

import (
	"math/rand"
	"time"

	"github.com/JuanObandoDeveloper/poker/internal/models"
)

var (
	Ranks = []models.Rank{models.Two, models.Three, models.Four, models.Five, models.Six, models.Seven, models.Eight, models.Nine, models.Ten, models.Jack, models.Queen, models.King, models.Ace}
	suits = []models.Suit{models.Spades, models.Hearts, models.Diamonds, models.Clubs}
)

type Poker struct {
	Deck []models.Card
}

func NewPoker() *Poker {
	return &Poker{}
}

func (p *Poker) Insert() {
	for _, suit := range suits {
		for _, rank := range Ranks {
			p.Deck = append(p.Deck, models.Card{Rank: rank, Suit: suit})
		}
	}
}

func (p *Poker) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(p.Deck), func(i, j int) { p.Deck[i], p.Deck[j] = p.Deck[j], p.Deck[i] })
}

func (p *Poker) Distribute() []models.Card {
	prevCards := []models.Card{}
	for i := 0; i < 5; i++ {
		card := p.GetCard()
		prevCards = append(prevCards, card)
	}
	return prevCards
}

func (p *Poker) GetCard() models.Card {
	index := p.RandomNum()
	card := p.Deck[index]
	p.Deck = append(p.Deck[:index], p.Deck[index+1:]...)
	return card
}

func (p *Poker) RandomNum() int {
	return rand.Intn(len(p.Deck))
}

func IsRoyalFlush(hand []models.Card) bool {
	for i := 0; i < len(hand)-4; i++ {
		if hand[i].Rank == models.Ten && hand[i+1].Rank == models.Jack && hand[i+2].Rank == models.Queen && hand[i+3].Rank == models.King && hand[i+4].Rank == models.Ace {
			if IsFlush(hand) {
				return true
			}
		}
	}
	return false
}

func IsStraightFlush(hand []models.Card) bool {
	if IsFlush(hand) && IsStraight(hand) {
		return true
	}
	return false
}

func IsPoker(hand []models.Card) bool {
	for i := 0; i < len(hand)-3; i++ {
		if hand[i].Rank == hand[i+1].Rank && hand[i+1].Rank == hand[i+2].Rank && hand[i+2].Rank == hand[i+3].Rank {
			return true
		}
	}
	return false
}

func IsFullHouse(hand []models.Card) bool {
	for i := 0; i < len(hand)-2; i++ {
		if hand[i].Rank == hand[i+1].Rank && hand[i+1].Rank == hand[i+2].Rank {
			for j := i + 3; j < len(hand)-1; j++ {
				if hand[j].Rank == hand[j+1].Rank {
					return true
				}
			}
		}
	}
	return false
}

func IsFlush(hand []models.Card) bool {
	suit := hand[0].Suit
	for _, card := range hand {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func IsStraight(hand []models.Card) bool {
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].Rank+1 != hand[i+1].Rank {
			return false
		}
	}
	return true
}

func IsThreeOfAKind(hand []models.Card) bool {
	for i := 0; i < len(hand)-2; i++ {
		if hand[i].Rank == hand[i+1].Rank && hand[i].Rank == hand[i+2].Rank {
			return true
		}
	}
	return false
}

func IsTwoPair(hand []models.Card) bool {
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].Rank == hand[i+1].Rank {
			for j := i + 2; j < len(hand)-1; j++ {
				if hand[j].Rank == hand[j+1].Rank {
					return true
				}
			}
		}
	}
	return false
}

func IsOnePair(hand []models.Card) bool {
	for i := 0; i < len(hand)-1; i++ {
		if hand[i].Rank == hand[i+1].Rank {
			if (hand[i].Rank == models.Jack && hand[i+1].Rank == models.Jack) ||
				(hand[i].Rank == models.Queen && hand[i+1].Rank == models.Queen) ||
				(hand[i].Rank == models.King && hand[i+1].Rank == models.King) ||
				(hand[i].Rank == models.Ace && hand[i+1].Rank == models.Ace) {
				return true
			}
		}
	}
	return false
}
