package domain

type SinglePlayHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewSinglePlayHandler(next *IMakeCardPatternHandler) *SinglePlayHandler {
	return &SinglePlayHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *SinglePlayHandler) match(cards []*Card) bool {
	return len(cards) == 1
}

func (handler *SinglePlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern) {
	return NewSingleCardPattern(cards)
}
