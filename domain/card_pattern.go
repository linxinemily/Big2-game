package domain

import (
	"big2/domain/enum"
	"fmt"
	"strings"
)

type CardPattern interface {
	isBiggerThan(another CardPattern) bool
	String() string
	containsClub3() bool
	getCards() []*Card
}

type AbstractCardPattern struct {
	cards []*Card
}

func NewAbstractCardPattern(cards []*Card) *AbstractCardPattern {
	return &AbstractCardPattern{
		cards: cards,
	}
}

func (cardPattern *AbstractCardPattern) String() string {
	var sb strings.Builder
	for _, card := range cardPattern.cards {
		sb.WriteString(fmt.Sprintf("rank: %s, suit: %s\n", card.Rank, card.Suit))
	}
	fmt.Println(fmt.Sprintf(sb.String()))
	return fmt.Sprintf(sb.String())
}

func (cardPattern *AbstractCardPattern) containsClub3() bool {
	for _, card := range cardPattern.cards {
		if card.Rank == enum.Three && card.Suit == enum.Club {
			return true
		}
	}

	return false
}

func (cardPattern *AbstractCardPattern) getCards() []*Card {
	return cardPattern.cards
}
