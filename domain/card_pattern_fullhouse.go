package domain

type FullHouseCardPattern struct{}

func (pattern *FullHouseCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*FullHouseCardPattern)
	if !ok {
		return false
	}

	return false
}
