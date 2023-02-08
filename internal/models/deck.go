package models

import (
	"math/rand"
	"sort"
	"time"
)

type Deck []Card

func NewMainDeck() *Deck {
	deck := Deck{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck = append(deck, Card{suit, rank})
		}
	}
	return &deck
}

func (d *Deck) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

func (d *Deck) DealDeck(size uint8) *Deck {
	deckToDeal := (*d)[:size]
	*d = (*d)[size:]
	return &deckToDeal
}

func (d *Deck) AddDeck(deckToAdd *Deck) {
	*d = append(*d, *deckToAdd...)
}

func (d *Deck) RemoveDeck(deckToRemove *Deck) {
	for _, card := range *deckToRemove {
		for i, deckCard := range *d {
			if card == deckCard {
				*d = append((*d)[:i], (*d)[i+1:]...)
				break
			}
		}
	}
}

func (d *Deck) Deck(indexes ...uint8) *Deck {
	deckToReturn := new(Deck)
	for _, index := range indexes {
		*deckToReturn = append(*deckToReturn, (*d)[index])
	}
	return deckToReturn
}

func (d *Deck) Sort() {
	sort.Slice(*d, func(i, j int) bool {
		return (*d)[i].Rank < (*d)[j].Rank
	})
}

func (d Deck) String() string {
	str := ""
	for _, c := range d {
		str += c.String() + " "
	}
	return str
}
