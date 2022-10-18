package domain

type StraightCardPattern struct{}

func (pattern *StraightCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*StraightCardPattern)
	if !ok {
		return false
	}

	return false
}
