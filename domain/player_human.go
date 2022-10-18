package domain

type HumanPlayer struct {
	*AbstractPlayer
}

func NewHumanPlayer(id int, playHandler *IPlayHandler) *HumanPlayer {
	return &HumanPlayer{
		AbstractPlayer: NewAbstractPlayer(id, playHandler),
	}
}

func (p *HumanPlayer) getCardsFromUserInput() []*Card {
	return []*Card{}
}
