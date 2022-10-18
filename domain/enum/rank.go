package enum

type Rank int

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Night
	Ten
	J
	Q
	K
	A
)

func (r Rank) String() string {
	return map[Rank]string{
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Night: "9",
		Ten:   "10",
		J:     "J",
		Q:     "Q",
		K:     "K",
		A:     "A",
	}[r]
}
