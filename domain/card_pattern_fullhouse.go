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
	pattern2, ok := another.(*FullHouseCardPattern)
	if !ok {
		return false
	}

	return pattern.getCards()[0].isBiggerThan(pattern2.getCards()[0])
}
