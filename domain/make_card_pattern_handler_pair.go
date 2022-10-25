package domain

import "sort"

type PairPlayHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewPairPlayHandler(next *IMakeCardPatternHandler) *PairPlayHandler {
	return &PairPlayHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *PairPlayHandler) match(cards []*Card) (cardPattern CardPattern, ok bool) {
	ok = len(cards) == 2 && cards[0].Rank == cards[1].Rank
	sort.Sort(sort.Reverse(CardSlice(cards))) //降冪排序
	return NewPairCardPattern(cards), ok
}
