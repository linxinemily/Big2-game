package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HumanPlayer struct {
	*AbstractPlayer
}

func NewHumanPlayer(id int, playHandler *IMakeCardPatternHandler) *HumanPlayer {
	return &HumanPlayer{
		AbstractPlayer: NewAbstractPlayer(id, playHandler),
	}
}

func (p *HumanPlayer) getCardsFromUserInput() []*Card {
	var res []*Card
	for {
		if res != nil {
			break
		}

		p.printHand()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		indexesOfCards := strings.Split(scanner.Text(), " ")

		if indexesOfCards[0] == "-1" {
			fmt.Println("-1 break")
			break
		}

		var cards []*Card
		for i := 0; i < len(indexesOfCards); i++ {
			intVar, err := strconv.Atoi(indexesOfCards[i])
			if err != nil {
				continue
			}
			cards = append(cards, p.hand[intVar])
		}
		res = cards
	}

	fmt.Println("end")
	return res
}

func (p *HumanPlayer) nameSelf() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter player name:")
	scanner.Scan()
	name := scanner.Text()
	p.name = name
}

func (p *HumanPlayer) printHand() {
	fmt.Println("Your hand:")
	for i, c := range p.hand {
		fmt.Printf("[%d] rank %s, suit %s \n", i, c.Rank, c.Suit)
	}
	fmt.Println("or enter [-1] to pass this turn")
}
