package domain

import (
	"big2/domain/enum"
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchStraight(t *testing.T) {
	validInputs := [][]*Card{
		{
			{Rank: enum.Three, Suit: enum.Heart},
			{Rank: enum.Four, Suit: enum.Diamond},
			{Rank: enum.Five, Suit: enum.Diamond},
			{Rank: enum.Six, Suit: enum.Spade},
			{Rank: enum.Seven, Suit: enum.Heart},
		},
		{
			{Rank: enum.Ten, Suit: enum.Heart},
			{Rank: enum.J, Suit: enum.Diamond},
			{Rank: enum.Q, Suit: enum.Diamond},
			{Rank: enum.K, Suit: enum.Spade},
			{Rank: enum.A, Suit: enum.Heart},
		},
		{
			{Rank: enum.J, Suit: enum.Heart},
			{Rank: enum.Q, Suit: enum.Diamond},
			{Rank: enum.K, Suit: enum.Diamond},
			{Rank: enum.A, Suit: enum.Spade},
			{Rank: enum.Two, Suit: enum.Heart},
		},
		{
			{Rank: enum.Two, Suit: enum.Diamond},
			{Rank: enum.Three, Suit: enum.Spade},
			{Rank: enum.Four, Suit: enum.Heart},
			{Rank: enum.Five, Suit: enum.Heart},
			{Rank: enum.Six, Suit: enum.Diamond},
		},
		{
			{Rank: enum.Two, Suit: enum.Diamond},
			{Rank: enum.Three, Suit: enum.Spade},
			{Rank: enum.Q, Suit: enum.Heart},
			{Rank: enum.K, Suit: enum.Heart},
			{Rank: enum.A, Suit: enum.Diamond},
		},
	}

	for i := 0; i < len(validInputs); i++ {
		t.Run(fmt.Sprintf("valid inputs %v", i), func(t *testing.T) {
			assert.True(t, NewStraightPlayHandler(nil).match(validInputs[i]))
		})
	}

	invalidInputs := [][]*Card{
		{
			{Rank: enum.Three, Suit: enum.Heart},
			{Rank: enum.A, Suit: enum.Diamond},
			{Rank: enum.Night, Suit: enum.Diamond},
			{Rank: enum.Night, Suit: enum.Spade},
			{Rank: enum.A, Suit: enum.Heart},
		},
		{
			{Rank: enum.Two, Suit: enum.Diamond},
			{Rank: enum.Five, Suit: enum.Spade},
			{Rank: enum.Six, Suit: enum.Heart},
			{Rank: enum.J, Suit: enum.Heart},
			{Rank: enum.A, Suit: enum.Diamond},
		},
	}

	for i := 0; i < len(invalidInputs); i++ {
		t.Run(fmt.Sprintf("invalid inputs %v", i), func(t *testing.T) {
			assert.False(t, NewStraightPlayHandler(nil).match(invalidInputs[i]))
		})
	}
}
