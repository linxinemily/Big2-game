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
	pattern2, ok := another.(*SingleCardPattern)
	if !ok {
		return false
	}

	return pattern.getCards()[0].isBiggerThan(pattern2.getCards()[0])
}
