package domain

import "big2/domain/enum"

type Card struct {
	Rank enum.Rank
	Suit enum.Suit
}

func NewCard(rank enum.Rank, suit enum.Suit) (c *Card) {
	return &Card{
		Rank: rank,
		Suit: suit,
	}
}

func (card *Card) isEuqalTo(card2 *Card) bool {
	return card.Rank == card2.Rank && card.Suit == card2.Suit
}
