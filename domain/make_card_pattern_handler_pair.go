package domain

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
	return NewPairCardPattern(cards), ok
}
