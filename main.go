package main

import "big2/domain"

func main() {

	playhandler := domain.NewIMakeCardPatternHandler(domain.NewSinglePlayHandler(
		domain.NewIMakeCardPatternHandler(domain.NewPairPlayHandler(
			domain.NewIMakeCardPatternHandler(domain.NewStraightPlayHandler(
				domain.NewIMakeCardPatternHandler(domain.NewFullHousePlayHandler(nil)),
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
