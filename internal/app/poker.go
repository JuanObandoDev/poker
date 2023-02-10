package app

import (
	"github.com/JuanObandoDeveloper/poker/internal/models"
)

type Poker struct {
	mainDeck   *models.Deck
	PlayerDeck *models.Deck
}

func NewPoker() *Poker {
	poker := new(Poker)
	poker.mainDeck = models.NewMainDeck()
	poker.mainDeck.Shuffle()
	poker.PlayerDeck = new(models.Deck)
	return poker
}

func (p *Poker) Distribute() {
	p.PlayerDeck = p.mainDeck.DealDeck(5)
}

func (p *Poker) DiscardDeck(deckToDiscard *models.Deck) {
	p.PlayerDeck.RemoveDeck(deckToDiscard)
	p.PlayerDeck.AddDeck(p.mainDeck.DealDeck(uint8(len(*deckToDiscard))))
}

func (p *Poker) IsRoyalFlush() bool {
	for i := 0; i < len(*p.PlayerDeck)-4; i++ {
		if (*p.PlayerDeck)[i].Rank == models.Ten && (*p.PlayerDeck)[i+1].Rank == models.Jack && (*p.PlayerDeck)[i+2].Rank == models.Queen && (*p.PlayerDeck)[i+3].Rank == models.King && (*p.PlayerDeck)[i+4].Rank == models.Ace {
			if p.IsFlush() {
				return true
			}
		}
	}
	return false
}

func (p *Poker) IsStraightFlush() bool {
	if p.IsFlush() && p.IsStraight() {
		return true
	}
	return false
}

func (p *Poker) IsPoker() bool {
	for i := 0; i < len(*p.PlayerDeck)-3; i++ {
		if (*p.PlayerDeck)[i].Rank == (*p.PlayerDeck)[i+1].Rank && (*p.PlayerDeck)[i+1].Rank == (*p.PlayerDeck)[i+2].Rank && (*p.PlayerDeck)[i+2].Rank == (*p.PlayerDeck)[i+3].Rank {
			return true
		}
	}
	return false
}

func (p *Poker) IsFullHouse() bool {
	for i := 0; i < len(*p.PlayerDeck)-2; i++ {
		if (*p.PlayerDeck)[i].Rank == (*p.PlayerDeck)[i+1].Rank && (*p.PlayerDeck)[i+1].Rank == (*p.PlayerDeck)[i+2].Rank {
			for j := i + 3; j < len((*p.PlayerDeck))-1; j++ {
				if (*p.PlayerDeck)[j].Rank == (*p.PlayerDeck)[j+1].Rank {
					return true
				}
			}
		}
	}
	return false
}

func (p *Poker) IsFlush() bool {
	suit := (*p.PlayerDeck)[0].Suit
	for _, card := range *p.PlayerDeck {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func (p *Poker) IsStraight() bool {
	for i := 0; i < len((*p.PlayerDeck))-1; i++ {
		if (*p.PlayerDeck)[i].Rank+1 != (*p.PlayerDeck)[i+1].Rank {
			return false
		}
	}
	return true
}

func (p *Poker) IsThreeOfAKind() bool {
	for i := 0; i < len(*p.PlayerDeck)-2; i++ {
		if (*p.PlayerDeck)[i].Rank == (*p.PlayerDeck)[i+1].Rank && (*p.PlayerDeck)[i].Rank == (*p.PlayerDeck)[i+2].Rank {
			return true
		}
	}
	return false
}

func (p *Poker) IsTwoPair() bool {
	for i := 0; i < len(*p.PlayerDeck)-1; i++ {
		if (*p.PlayerDeck)[i].Rank == (*p.PlayerDeck)[i+1].Rank {
			for j := i + 2; j < len((*p.PlayerDeck))-1; j++ {
				if (*p.PlayerDeck)[j].Rank == (*p.PlayerDeck)[j+1].Rank {
					return true
				}
			}
		}
	}
	return false
}

func (p *Poker) IsOnePair() bool {
	for i := 0; i < len(*p.PlayerDeck)-1; i++ {
		if (*p.PlayerDeck)[i].Rank == (*p.PlayerDeck)[i+1].Rank {
			if ((*p.PlayerDeck)[i].Rank == models.Jack && (*p.PlayerDeck)[i+1].Rank == models.Jack) ||
				((*p.PlayerDeck)[i].Rank == models.Queen && (*p.PlayerDeck)[i+1].Rank == models.Queen) ||
				((*p.PlayerDeck)[i].Rank == models.King && (*p.PlayerDeck)[i+1].Rank == models.King) ||
				((*p.PlayerDeck)[i].Rank == models.Ace && (*p.PlayerDeck)[i+1].Rank == models.Ace) {
				return true
			}
		}
	}
	return false
}

func (p *Poker) Validate() bool {
	p.PlayerDeck.Sort()
	if p.IsRoyalFlush() {
		return true
	}
	if p.IsStraightFlush() {
		return true
	}
	if p.IsPoker() {
		return true
	}
	if p.IsFullHouse() {
		return true
	}
	if p.IsFlush() {
		return true
	}
	if p.IsStraight() {
		return true
	}
	if p.IsThreeOfAKind() {
		return true
	}
	if p.IsTwoPair() {
		return true
	}
	if p.IsOnePair() {
		return true
	}
	return false
}
