package domain

import "fmt"

type Big2 struct {
	round       int
	topPlay     CardPattern
	topPlayer   *IPlayer
	winner      *IPlayer
	players     []*IPlayer
	deck        *Deck
	playHandler *IPlayHandler
}

func NewBig2(playHandler *IPlayHandler) *Big2 {
	return &Big2{
		round:       1,
		deck:        NewDeck(),
		playHandler: playHandler,
	}
}

func (big2 *Big2) GenerateHumanPlayer() {
	player := &IPlayer{
		Player: NewHumanPlayer(len(big2.players), big2.playHandler),
	}
	big2.players = append(big2.players, player)
}

func (big2 *Big2) GenerateAIPlayer() {
	player := &IPlayer{
		Player: NewAIPlayer(len(big2.players), big2.playHandler),
	}
	big2.players = append(big2.players, player)
}

func (big2 *Big2) playerDrawCards() {
	count := 0
	for {
		if len(big2.deck.getCards()) == 0 {
			break
		}
		deck := big2.deck
		card := deck.deal()
		p := big2.players[count%4]
		p.addCardIntoHand(card)
		count += 1
	}

}

func (big2 *Big2) Start() {

	big2.deck.shuffle()

	big2.playerDrawCards()

	for big2.winner == nil {
		big2.takeRound()
	}
}

func (big2 *Big2) takeRound() {

	fmt.Printf("----Round %v----\n", big2.round)

	var player *IPlayer

	if big2.round == 1 { // 第一回合
		player = big2.getPlayerHasClub3()
	} else {
		player = big2.topPlayer
	}

	for turn := 0; turn < len(big2.players); turn++ {

		cardPattern := player.takeTurn(big2.topPlay, turn)

		if len(player.Hand()) == 0 { // 打完牌後檢查玩家還有沒有剩餘的牌
			big2.winner = player
			break
		}

		if cardPattern != nil {
			big2.topPlay = cardPattern
			big2.topPlayer = player
		}

		nextPlayerId := (player.Id() + 1) / len(big2.players)
		player = big2.players[nextPlayerId]
	}

	big2.topPlay = nil

}

func (big2 *Big2) ShowWinner() *IPlayer {
	return big2.winner
}

func (big2 *Big2) getPlayerHasClub3() *IPlayer {
	return &IPlayer{}
}
