package domain

type FullHouseCardPattern struct {
	*AbstractCardPattern
}

func NewFullHouseCardPattern(cards []*Card) *FullHouseCardPattern {
	return &FullHouseCardPattern{
		AbstractCardPattern: NewAbstractCardPattern(cards),
	}
}

func (pattern *FullHouseCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*FullHouseCardPattern)
	if !ok {
		return false
	}

	return false
}
