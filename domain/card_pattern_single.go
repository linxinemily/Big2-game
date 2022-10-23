package domain

type SingleCardPattern struct {
	*AbstractCardPattern
}

func NewSingleCardPattern(cards []*Card) *SingleCardPattern {
	return &SingleCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *SingleCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*SingleCardPattern)
	if !ok {
		return false
	}

	return false
}
