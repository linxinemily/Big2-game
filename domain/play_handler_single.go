package domain

type SinglePlayHandler struct {
	*AbstractPlayHandler
}

func NewSinglePlayHandler(next *IPlayHandler) *SinglePlayHandler {
	return &SinglePlayHandler{
		AbstractPlayHandler: NewAbstractHandler(next),
	}
}

func (handler *SinglePlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern, makeSuccessfully bool) {
	return nil, false
}
