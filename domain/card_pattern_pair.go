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
	pattern2, ok := another.(*PairCardPattern)
	if !ok {
		return false
	}

	return pattern.getCards()[0].isBiggerThan(pattern2.getCards()[0])
}
