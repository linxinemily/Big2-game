package domain

type SinglePlayHandler struct {
	*AbstractMakeCardPatternHandler
}

func NewSinglePlayHandler(next *IMakeCardPatternHandler) *SinglePlayHandler {
	return &SinglePlayHandler{
		AbstractMakeCardPatternHandler: NewAbstractHandler(next),
	}
}

func (handler *SinglePlayHandler) match(cards []*Card) (cardPattern CardPattern, ok bool) {
	ok = len(cards) == 1
	return NewSingleCardPattern(cards), ok
}
