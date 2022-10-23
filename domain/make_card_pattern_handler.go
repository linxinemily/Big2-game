package domain

type IMakeCardPatternHandler struct {
	MakeCardPatternHandler
}

func NewIMakeCardPatternHandler(handler MakeCardPatternHandler) *IMakeCardPatternHandler {
	return &IMakeCardPatternHandler{
		MakeCardPatternHandler: handler,
	}
}

func (handler *IMakeCardPatternHandler) handle(cards []*Card) CardPattern { // template method

	if handler.match(cards) {
		return handler.makeCardPattern(cards)
	} else if handler.getNext() != nil {
		return handler.getNext().handle(cards)
	}

	return nil
}

type MakeCardPatternHandler interface {
	match(cards []*Card) bool
	makeCardPattern(cards []*Card) (cardPattern CardPattern)
	getNext() *IMakeCardPatternHandler
}

type AbstractMakeCardPatternHandler struct {
	cardPatternType string
	next            *IMakeCardPatternHandler
}

func NewAbstractHandler(next *IMakeCardPatternHandler) *AbstractMakeCardPatternHandler {
	return &AbstractMakeCardPatternHandler{
		next: next,
	}
}

func (handler *AbstractMakeCardPatternHandler) getNext() *IMakeCardPatternHandler {
	return handler.next
}
