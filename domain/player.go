package domain

import "fmt"

type IPlayer struct {
	Player
}

func (p *IPlayer) takeTurn(topPlay CardPattern, turn int) CardPattern {
	fmt.Println(p)
	fmt.Printf("It's player %s's turn\n", p.Name())
	canPass := true
	if turn == 0 { // 新回合首輪玩家不能喊 pass
		canPass = false
	}

	for {
		cards := p.getCardsFromUserInput()
		if len(cards) == 0 { // pass
			if canPass {

				fmt.Printf("player %s pass\n", p.Name())
				return nil
			} else {
				fmt.Println("cannot pass")
			}
		} else {
			cardPattern := p.play(cards, topPlay)
			if cardPattern != nil {
				return cardPattern
			}
		}
	}
}

type Player interface {
	// pass()
	play(cards []*Card, topCardPattern CardPattern) CardPattern
	nameSelf()
	addCardIntoHand(card *Card)
	Hand() []*Card
	Id() int
	Name() string
	getCardsFromUserInput() []*Card
}

type AbstractPlayer struct {
	PlayHandler *IPlayHandler
	name        string
	id          int
	hand        []*Card
}

func NewAbstractPlayer(id int, playHandler *IPlayHandler) *AbstractPlayer {
	return &AbstractPlayer{
		id:          id,
		PlayHandler: playHandler,
	}
}

// func (p *AbstractPlayer) pass(input int) bool {
// 	return input == -1
// }

func (p *AbstractPlayer) play(cards []*Card, topCardPattern CardPattern) CardPattern {
	return p.PlayHandler.play(cards, topCardPattern)
}

func (p *AbstractPlayer) nameSelf() {

}

func (p *AbstractPlayer) addCardIntoHand(card *Card) {
	p.hand = append(p.hand, card)
}

func (p *AbstractPlayer) Hand() []*Card {
	return p.hand
}

func (p *AbstractPlayer) Id() int {
	return p.id
}

func (p *AbstractPlayer) Name() string {
	return p.name
}
