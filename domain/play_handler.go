package domain

import "reflect"

type IPlayHandler struct {
	PlayHandler
}

func NewIPlayHandler(handler PlayHandler) *IPlayHandler {
	return &IPlayHandler{
		PlayHandler: handler,
	}
}

func (handler *IPlayHandler) play(cards []*Card, topCardPattern CardPattern) CardPattern { // template method
	cardPattern, makeSuccessfully := handler.makeCardPattern(cards) // 在裡面 validate, 嘗試建出特定 CardPattern

	// case 1: 輸入的牌型合法，且有頂牌的狀況下，牌型和頂牌一樣，且比頂牌大，可出牌
	if makeSuccessfully && topCardPattern != nil && reflect.TypeOf(topCardPattern).String() == handler.getCardPatternType() && cardPattern.isBiggerThan(topCardPattern) {
		return cardPattern
	} else if makeSuccessfully && topCardPattern == nil { // case 2: 輸入牌型合法，無頂牌狀況下，可出牌
		return cardPattern
	} else if handler.getNext() != nil { // case 3: 皆非以上兩者情況，交給下個 handler 處理
		handler.getNext().play(cards, topCardPattern)
	}

	return nil // case 4: 以上情況皆非且已經沒有下個 handler 可處理
}

type PlayHandler interface {
	makeCardPattern(cards []*Card) (cardPattern CardPattern, makeSuccessfully bool)
	getCardPatternType() string
	getNext() *IPlayHandler
}

type AbstractPlayHandler struct {
	cardPatternType string
	next            *IPlayHandler
}

func NewAbstractHandler(next *IPlayHandler) *AbstractPlayHandler {
	return &AbstractPlayHandler{
		next: next,
	}
}

func (handler *AbstractPlayHandler) getCardPatternType() string {
	return ""
}

func (handler *AbstractPlayHandler) getNext() *IPlayHandler {
	return handler.next
}