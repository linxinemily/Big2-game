package domain

type PairPlayHandler struct {
	*AbstractPlayHandler
}

func NewPairPlayHandler(next *IPlayHandler) *PairPlayHandler {
	return &PairPlayHandler{
		AbstractPlayHandler: NewAbstractHandler(next),
	}
}

func (handler *PairPlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern, makeSuccessfully bool) {
	return nil, false
}
