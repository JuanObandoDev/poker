package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JuanObandoDeveloper/poker/internal/app"
)

func main() {
	poker := app.NewPoker()
	poker.Distribute()

	fmt.Println("your deck:", poker.PlayerDeck)

	fmt.Println("if you want to discard some cards, enter the index of the cards separated by spaces (e.g. 1 2 3 4 5) or enter to continue")
	indexes := indexesToDiscard()

	deckToDiscard := poker.PlayerDeck.Deck(indexes...)
	poker.DiscardDeck(deckToDiscard)

	fmt.Println("your deck:", poker.PlayerDeck)

	win := poker.Validate()
	if win {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose!")
	}
}

func indexesToDiscard() []uint8 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	indexes := numbers(scanner.Text())
	return indexes
}

func numbers(s string) []uint8 {
	var n []uint8
	for _, r := range strings.Fields(s) {
		i, err := strconv.Atoi(r)

		if err == nil {
			n = append(n, uint8(i-1))
		}
	}
	return n
}
