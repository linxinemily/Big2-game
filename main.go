package main

import "big2/domain"

func main() {

	playhandler := domain.NewIPlayHandler(domain.NewSinglePlayHandler(
		domain.NewIPlayHandler(domain.NewPairPlayHandler(
			domain.NewIPlayHandler(domain.NewStraightPlayHandler(
				domain.NewIPlayHandler(domain.NewFullHousePlayHandler(nil)),
			)),
		)),
	))

	big2 := domain.NewBig2(playhandler)

	big2.GenerateHumanPlayer()
	big2.GenerateAIPlayer()
	big2.GenerateAIPlayer()
	big2.GenerateAIPlayer()

	big2.Start()
	big2.ShowWinner()

}
