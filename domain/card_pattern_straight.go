package domain

type StraightCardPattern struct {
	*AbstractCardPattern
}

func NewStraightCardPattern(cards []*Card) *StraightCardPattern {
	return &StraightCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *StraightCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*StraightCardPattern)
	if !ok {
		return false
	}

	return false
}
