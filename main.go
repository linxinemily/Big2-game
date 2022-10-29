package main

import "big2/domain"

func main() {

	makeCardPatternHandler := domain.NewIMakeCardPatternHandler(domain.NewMakeSingleCardPatternHandler(
		domain.NewIMakeCardPatternHandler(domain.NewMakePairCardPatternHandler(
			domain.NewIMakeCardPatternHandler(domain.NewMakeStraightCardPatternHandler(
				domain.NewIMakeCardPatternHandler(domain.NewMakeFullHouseCardPatternHandler(nil)),
			)),
		)),
	))

	big2 := domain.NewBig2(makeCardPatternHandler)

	big2.GenerateHumanPlayer()
	big2.GenerateHumanPlayer()
	big2.GenerateHumanPlayer()
	big2.GenerateHumanPlayer()

	big2.Start()
	big2.ShowWinner()

}
