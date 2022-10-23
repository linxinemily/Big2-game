package domain

import (
	"big2/domain/enum"
	"fmt"
	"reflect"
)

type Big2 struct {
	round       int
	topPlay     CardPattern
	topPlayer   *IPlayer
	winner      *IPlayer
	players     []*IPlayer
	deck        *Deck
	playHandler *IMakeCardPatternHandler
}

func NewBig2(playHandler *IMakeCardPatternHandler) *Big2 {
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
	player.nameSelf()
}

func (big2 *Big2) GenerateAIPlayer() {
	player := &IPlayer{
		Player: NewAIPlayer(len(big2.players), big2.playHandler),
	}
	big2.players = append(big2.players, player)
	player.nameSelf()
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
		//player = big2.getPlayerHasClub3()
		player = big2.players[0]
		fmt.Printf("player %s has Club3\n", player.Name())
	} else {
		player = big2.topPlayer
		fmt.Printf("top player is %s\n", player.Name())
	}

	for turn := 0; turn < len(big2.players); turn++ {

		var cardPattern CardPattern

		for {
			cardPattern = player.takeTurn(turn)

			//if big2.round == 1 && turn == 0 && !cardPattern.containsClub3() { // 第一回合首輪玩家只能出包含梅花3的牌型
			//	fmt.Println("出的牌型不包含梅花3")
			//	continue
			//}

			// 如果有出牌(沒有 pass)，在有頂牌的狀況下，牌型必須和頂牌一樣，且比頂牌大
			if cardPattern != nil && big2.topPlay != nil {
				if reflect.TypeOf(big2.topPlay).String() != reflect.TypeOf(cardPattern).String() {
					fmt.Printf("牌型和頂牌不一樣，不可出該副牌, top Play: %s\n", big2.topPlay)
					continue
				}
				if !cardPattern.isBiggerThan(big2.topPlay) {
					fmt.Printf("沒有比頂牌大，不可出該副牌, top Play: %s\n", big2.topPlay)
					continue
				}
			}
			break
		}

		if cardPattern != nil {
			for _, card := range cardPattern.getCards() {
				player.removeCardFromHand(card)
			}
		}

		if len(player.Hand()) == 0 { // 打完牌後檢查玩家還有沒有剩餘的牌
			big2.winner = player
			break
		}

		if cardPattern != nil {
			big2.topPlay = cardPattern
			big2.topPlayer = player
			fmt.Printf("top Play: %s\n", big2.topPlay.String())
		}

		nextPlayerId := (player.Id() + 1) % len(big2.players)
		player = big2.players[nextPlayerId]
	}

	big2.topPlay = nil
	big2.round += 1

}

func (big2 *Big2) ShowWinner() *IPlayer {
	return big2.winner
}

func (big2 *Big2) getPlayerHasClub3() *IPlayer {
	for _, player := range big2.players {
		for _, card := range player.Hand() {
			if card.Suit == enum.Club && card.Rank == enum.Three {
				return player
			}
		}
	}
	return nil
}
