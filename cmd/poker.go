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

	fmt.Println("if you want to keep some cards, enter the index of the cards separated by spaces (e.g. 1 2 3 4 5) or enter to discard all cards")
	indexes := indexesToKeep()
	indexesToDiscard := keepToDiscard(indexes)

	deckToDiscard := poker.PlayerDeck.Deck(indexesToDiscard...)
	poker.DiscardDeck(deckToDiscard)

	fmt.Println("your deck:", poker.PlayerDeck)

	win := poker.Validate()
	if win {
		fmt.Println("You win!")
	} else {
		fmt.Println("You lose!")
	}
}

func indexesToKeep() []uint8 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	indexes := numbers(scanner.Text())
	return indexes
}

func keepToDiscard(indexes []uint8) []uint8 {
	indexesToDiscard := []uint8{0, 1, 2, 3, 4}
	for _, index := range indexes {
		for i, indexToDiscard := range indexesToDiscard {
			if index == indexToDiscard {
				indexesToDiscard = append(indexesToDiscard[:i], indexesToDiscard[i+1:]...)
			}
		}
	}
	return indexesToDiscard
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
