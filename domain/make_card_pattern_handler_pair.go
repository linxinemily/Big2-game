package domain

type PairPlayHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewPairPlayHandler(next *IMakeCardPatternHandler) *PairPlayHandler {
	return &PairPlayHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *PairPlayHandler) match(cards []*Card) bool {
	return len(cards) == 2 && cards[0].Rank == cards[1].Rank
}

func (handler *PairPlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern) {
	return NewPairCardPattern(cards)
}
