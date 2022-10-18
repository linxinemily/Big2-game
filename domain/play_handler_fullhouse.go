package domain

type FullHousePlayHandler struct {
	*AbstractPlayHandler
}

func NewFullHousePlayHandler(next *IPlayHandler) *FullHousePlayHandler {
	return &FullHousePlayHandler{
		AbstractPlayHandler: NewAbstractHandler(next),
	}
}

func (handler *FullHousePlayHandler) makeCardPattern(cards []*Card) (cardPattern CardPattern, makeSuccessfully bool) {
	return nil, false
}
