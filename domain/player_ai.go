package domain

type AIPlayer struct {
	*AbstractPlayer
}

func NewAIPlayer(id int, playHandler *IPlayHandler) *AIPlayer {
	return &AIPlayer{
		AbstractPlayer: NewAbstractPlayer(id, playHandler),
	}
}

func (p *AIPlayer) getCardsFromUserInput() []*Card {
		return []*Card{}
}
