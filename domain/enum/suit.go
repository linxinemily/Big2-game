package enum

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	return map[Suit]string{
		Club:    "♣︎",
		Diamond: "♦︎",
		Heart:   "♥︎",
		Spade:   "♠️",
	}[s]
}
