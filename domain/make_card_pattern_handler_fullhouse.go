package domain

import (
	"big2/domain/enum"
	"sort"
)

type FullHousePlayHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewFullHousePlayHandler(next *IMakeCardPatternHandler) *FullHousePlayHandler {
	return &FullHousePlayHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *FullHousePlayHandler) match(cards []*Card) bool {
	if len(cards) != 5 {
		return false
	}
	// 1. group by rank
	m := make(map[enum.Rank]int)

	for _, card := range cards {
		if _, exists := m[card.Rank]; !exists {
			m[card.Rank] = 1
		} else {
			m[card.Rank] += 1
		}
	}

	type CalculatedResult struct {
		Rank  enum.Rank
		Count int
	}
	var group []CalculatedResult
	for key, val := range m {
		group = append(group, CalculatedResult{Rank: key, Count: val})
	}

	// 2. sort by count desc
	sort.Slice(group, func(i, j int) bool {
		return group[i].Count < group[j].Count
	})
	// 3. first group's count should be 3 and second group's count should be 2
	return group[0].Count == 2 && group[1].Count == 3
}

func (handler *FullHousePlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern) {
	return NewFullHouseCardPattern(cards)
}
