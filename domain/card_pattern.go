package domain

type CardPattern interface {
	isBiggerThan(another CardPattern) bool
}

