package domain

import (
	"big2/domain/enum"
	"sort"
)

type StraightPlayHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewStraightPlayHandler(next *IMakeCardPatternHandler) *StraightPlayHandler {
	return &StraightPlayHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *StraightPlayHandler) match(cards []*Card) bool {
	if len(cards) != 5 {
		return false
	}

	// 先依照大小（升冪）排序完，再分組
	// 分組規則：如果中間有斷層(相減大於一)就分成另外一組
	// 最後結果如果第一組的第一個是 Three(0)，且第二組的最後一個是 Two(12)，表示可以接起來 => 為合法牌型
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})

	group := []map[int]bool{
		{},
	}

	for i := 0; i < len(cards)-1; i++ {

		group[len(group)-1][int(cards[i].Rank)] = true

		if int(cards[i+1].Rank)-int(cards[i].Rank) == 1 {
			group[len(group)-1][int(cards[i+1].Rank)] = true
		} else {
			group = append(group, map[int]bool{
				int(cards[i+1].Rank): true,
			})
		}
	}

	if len(group[0]) == 5 {
		return true
	}

	if len(group[1]) > 0 {
		var firstGroupKeys []int
		for k := range group[0] {
			firstGroupKeys = append(firstGroupKeys, k)
		}
		sort.Ints(firstGroupKeys)

		var SecondGroupKeys []int
		for k := range group[1] {
			SecondGroupKeys = append(SecondGroupKeys, k)
		}
		sort.Ints(SecondGroupKeys)
		return firstGroupKeys[0] == int(enum.Three) && SecondGroupKeys[len(SecondGroupKeys)-1] == int(enum.Two)
	}

	return false
}

func (handler *StraightPlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern) {
	return NewStraightCardPattern(cards)
}
