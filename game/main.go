package main

import (
	"fmt"

	"github.com/JuanObandoDeveloper/poker/internal/app"
)

var (
	deck = app.NewPoker()
)

func main() {
	deck.Insert()
	deck.Shuffle()
	hand := deck.Distribute()
	app.ShowHand(hand)
	hand = app.KeepOrDiscard(hand, *deck)
	hand = app.Sort(hand)
	app.ShowHand(hand)
	res := app.Validate(hand)
	if res {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose!")
	}
}
