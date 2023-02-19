package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JuanObandoDeveloper/poker/internal/app"
)

var (
	balance    int     = 1000000
	bet        float64 = 0
	moneyToBet float64 = 83.333333333333333333
)

func main() {
	menu()
}

func menu() {
	var op int
	for op != 6 {
		op = 0
		fmt.Println("1. Play")
		fmt.Println("2. Rules")
		fmt.Println("3. Up Bets")
		fmt.Println("4. Down Bets")
		fmt.Println("5. Balance")
		fmt.Println("6. Exit")
		fmt.Print("Select an option: ")
		fmt.Scanln(&op)
		switch op {
		case 1:
			if bet == 0 {
				fmt.Println("You have to bet first")
				break
			}
			if balance <= 0 {
				fmt.Println("You have no money left")
				op = 6
				break
			}
			play()
		case 2:
			fmt.Println("The rules are simple, you have 5 cards and you have to choose which ones you want to keep and which ones you want to discard. You can bet up until 5 coins, and you can win with: Pair (only with Jacks or better), Two pairs, Three of a kind, Straight, Flush, Full house, Four of a kind, Straight flush, and Royal flush")
		case 3:
			if bet == 5 {
				fmt.Println("You can't bet more than 5 coins")
				break
			}
			bet++
			moneyToBet = moneyToBet * bet
			fmt.Printf("Your bet is now: %v coins and money to bet: $%v\n", bet, int(moneyToBet))
		case 4:
			if bet == 0 {
				fmt.Println("You can't bet less than 0 coins")
				break
			}
			moneyToBet = moneyToBet / bet
			bet--
			fmt.Printf("Your bet is now: %v coins and money to bet: $%v\n", bet, int(moneyToBet))
		case 5:
			fmt.Println("Your balance is: $", balance)
		case 6:
			fmt.Println("Bye!")
		default:
			fmt.Println("Invalid option")
		}
	}
}

func play() {
	poker := app.NewPoker()
	poker.Distribute()

	fmt.Println("your deck:", poker.PlayerDeck)

	fmt.Println("if you want to keep some cards, enter the index of the cards separated by spaces (e.g. 1 2 3 4 5) or enter to discard all cards")
	indexes := indexesToKeep()
	indexesToDiscard := keepToDiscard(indexes)

	deckToDiscard := poker.PlayerDeck.Deck(indexesToDiscard...)
	poker.DiscardDeck(deckToDiscard)

	fmt.Println("your deck:", poker.PlayerDeck)

	win, winHand := poker.Validate()
	if win {
		switch winHand {
		case 9:
			fmt.Println("You win! Royal flush")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*250))
				balance += int(moneyToBet * 250)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*500))
				balance += int(moneyToBet * 500)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*750))
				balance += int(moneyToBet * 750)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*1000))
				balance += int(moneyToBet * 1000)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*4000))
				balance += int(moneyToBet * 4000)
			}
			fmt.Println("Your balance:", balance)
		case 8:
			fmt.Println("You win! Straight flush")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*50))
				balance += int(moneyToBet * 50)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*100))
				balance += int(moneyToBet * 100)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*150))
				balance += int(moneyToBet * 150)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*200))
				balance += int(moneyToBet * 200)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*250))
				balance += int(moneyToBet * 250)
			}
			fmt.Println("Your balance:", balance)
		case 7:
			fmt.Println("You win! Four of a kind")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*25))
				balance += int(moneyToBet * 25)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*50))
				balance += int(moneyToBet * 50)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*75))
				balance += int(moneyToBet * 75)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*100))
				balance += int(moneyToBet * 100)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*125))
				balance += int(moneyToBet * 125)
			}
			fmt.Println("Your balance:", balance)
		case 6:
			fmt.Println("You win! Full house")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*9))
				balance += int(moneyToBet * 9)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*18))
				balance += int(moneyToBet * 18)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*27))
				balance += int(moneyToBet * 27)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*36))
				balance += int(moneyToBet * 36)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*45))
				balance += int(moneyToBet * 45)
			}
			fmt.Println("Your balance:", balance)
		case 5:
			fmt.Println("You win! Flush")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*6))
				balance += int(moneyToBet * 6)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*12))
				balance += int(moneyToBet * 12)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*18))
				balance += int(moneyToBet * 18)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*24))
				balance += int(moneyToBet * 24)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*30))
				balance += int(moneyToBet * 30)
			}
			fmt.Println("Your balance:", balance)
		case 4:
			fmt.Println("You win! Straight")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*4))
				balance += int(moneyToBet * 4)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*8))
				balance += int(moneyToBet * 8)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*12))
				balance += int(moneyToBet * 12)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*16))
				balance += int(moneyToBet * 16)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*20))
				balance += int(moneyToBet * 20)
			}
			fmt.Println("Your balance:", balance)
		case 3:
			fmt.Println("You win! Three of a kind")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*3))
				balance += int(moneyToBet * 3)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*6))
				balance += int(moneyToBet * 6)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*9))
				balance += int(moneyToBet * 9)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*12))
				balance += int(moneyToBet * 12)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*15))
				balance += int(moneyToBet * 15)
			}
			fmt.Println("Your balance:", balance)
		case 2:
			fmt.Println("You win! Two pairs")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet*2))
				balance += int(moneyToBet * 2)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*4))
				balance += int(moneyToBet * 4)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*6))
				balance += int(moneyToBet * 6)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*8))
				balance += int(moneyToBet * 8)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*10))
				balance += int(moneyToBet * 10)
			}
			fmt.Println("Your balance:", balance)
		case 1:
			fmt.Println("You win! Pair")
			switch bet {
			case 1:
				fmt.Println("You gain: ", int(moneyToBet))
				balance += int(moneyToBet)
			case 2:
				fmt.Println("You gain: ", int(moneyToBet*2))
				balance += int(moneyToBet * 2)
			case 3:
				fmt.Println("You gain: ", int(moneyToBet*3))
				balance += int(moneyToBet * 3)
			case 4:
				fmt.Println("You gain: ", int(moneyToBet*4))
				balance += int(moneyToBet * 4)
			case 5:
				fmt.Println("You gain: ", int(moneyToBet*5))
				balance += int(moneyToBet * 5)
			}
			fmt.Println("Your balance:", balance)
		}
	} else {
		fmt.Println("You lose!")
		balance -= int(moneyToBet * bet)
		fmt.Println("Your balance:", balance)
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
