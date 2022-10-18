package domain

type SingleCardPattern struct{}

func (pattern *SingleCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*SingleCardPattern)
	if !ok {
		return false
	}

	return false
}
