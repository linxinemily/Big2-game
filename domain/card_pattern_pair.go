package domain

type PairCardPattern struct {
	*AbstractCardPattern
}

func NewPairCardPattern(cards []*Card) *PairCardPattern {
	return &PairCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *PairCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*PairCardPattern)
	if !ok {
		return false
	}

	return false
}
