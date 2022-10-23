package domain

import (
	"big2/domain/enum"
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchFullHouse(t *testing.T) {
	validInputs := [][]*Card{
		{
			{Rank: enum.Night, Suit: enum.Heart},
			{Rank: enum.A, Suit: enum.Diamond},
			{Rank: enum.Night, Suit: enum.Diamond},
			{Rank: enum.Night, Suit: enum.Spade},
			{Rank: enum.A, Suit: enum.Heart},
		},
		{
			{Rank: enum.Five, Suit: enum.Heart},
			{Rank: enum.Five, Suit: enum.Diamond},
			{Rank: enum.Five, Suit: enum.Spade},
			{Rank: enum.Two, Suit: enum.Spade},
			{Rank: enum.Two, Suit: enum.Heart},
		},
	}

	for i := 0; i < len(validInputs); i++ {
		t.Run(fmt.Sprintf("valid inputs %v", i), func(t *testing.T) {
			assert.True(t, NewFullHousePlayHandler(nil).match(validInputs[i]))
		})
	}

	t.Run("invalid input", func(t *testing.T) {
		cards := []*Card{
			{
				Rank: enum.Three,
				Suit: enum.Heart,
			},
			{
				Rank: enum.A,
				Suit: enum.Diamond,
			},
			{
				Rank: enum.Night,
				Suit: enum.Diamond,
			},
			{
				Rank: enum.Night,
				Suit: enum.Spade,
			},
			{
				Rank: enum.A,
				Suit: enum.Heart,
			},
		}

		assert.False(t, NewFullHousePlayHandler(nil).match(cards))
	})
}
