package domain

type PairCardPattern struct{}

func (pattern *PairCardPattern) isBiggerThan(another CardPattern) bool {
	_, ok := another.(*PairCardPattern)
	if !ok {
		return false
	}

	return false
}
