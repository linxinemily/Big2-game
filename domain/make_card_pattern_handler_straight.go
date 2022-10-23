package domain

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

	for i := 0; i < len(cards)-1; i++ {
		if int(cards[i+1].Rank)-int(cards[i].Rank) != 1 && int(cards[i].Rank)-int(cards[i+1].Rank) != 12 {
			return false
		}
	}
	return true
}

func (handler *StraightPlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern) {
	return NewStraightCardPattern(cards)
}
