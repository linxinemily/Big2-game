package domain

type StraightPlayHandler struct {
	*AbstractPlayHandler
}

func NewStraightPlayHandler(next *IPlayHandler) *StraightPlayHandler {
	return &StraightPlayHandler{
		AbstractPlayHandler: NewAbstractHandler(next),
	}
}

func (handler *StraightPlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern, makeSuccessfully bool) {
	return nil, false
}
