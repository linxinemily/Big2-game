package main

import "big2/domain"

func main() {

	playhandler := domain.NewIMakeCardPatternHandler(domain.NewMakeSingleCardPatternHandler(
		domain.NewIMakeCardPatternHandler(domain.NewMakePairCardPatternHandler(
			domain.NewIMakeCardPatternHandler(domain.NewMakeStraightCardPatternHandler(
				domain.NewIMakeCardPatternHandler(domain.NewMakeFullHouseCardPatternHandler(nil)),
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
