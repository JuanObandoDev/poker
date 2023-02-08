package app

import (
	"fmt"
	"sort"

	"github.com/JuanObandoDeveloper/poker/internal/models"
)

func ShowHand(hand []models.Card) {
	fmt.Println("your hand: ", hand)
}

func KeepOrDiscard(hand []models.Card, deck Poker) []models.Card {
	keep := []models.Card{}
	for i := 0; i < len(hand); i++ {
		op := ConfirmKeep(hand, i)
		if op == 1 {
			keep = append(keep, hand[i])
		} else if op == 2 {
			keep = hand
			break
		} else if op == 3 {
			break
		} else if op == 0 {
			hand = append(hand[:i], hand[i+1:]...)
			i--
		}
	}
	for len(keep) < 5 {
		card := deck.GetCard()
		keep = append(keep, card)
	}
	return keep
}

func Sort(hand []models.Card) []models.Card {
	sort.Slice(hand, func(i, j int) bool {
		return hand[i].Value < hand[j].Value
	})
	return hand
}

func ConfirmKeep(hand []models.Card, index int) int {
	var op int
	fmt.Println("keep card ", hand[index], "? (1 for keep, enter for discard, 2 for keep all, 3 for discard all)")
	fmt.Scanln(&op)
	return op
}

func Validate(hand []models.Card) bool {
	var res bool
	if IsRoyalFlush(hand) {
		res = true
	} else if IsStraightFlush(hand) {
		res = true
	} else if IsPoker(hand) {
		res = true
	} else if IsFullHouse(hand) {
		res = true
	} else if IsFlush(hand) {
		res = true
	} else if IsStraight(hand) {
		res = true
	} else if IsThreeOfAKind(hand) {
		res = true
	} else if IsTwoPair(hand) {
		res = true
	} else if IsOnePair(hand) {
		res = true
	} else {
		res = false
	}
	return res
}
