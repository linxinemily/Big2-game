package domain

import "fmt"

type AIPlayer struct {
	*AbstractPlayer
}

func NewAIPlayer(id int, playHandler *IMakeCardPatternHandler) *AIPlayer {
	return &AIPlayer{
		AbstractPlayer: NewAbstractPlayer(id, playHandler),
	}
}

func (p *AIPlayer) getCardsFromUserInput() []*Card {
	return []*Card{}
}

func (p *AIPlayer) nameSelf() {
	p.name = fmt.Sprintf("AI Player %v", p.id)
}
