package main

import (
	"fmt"

	"github.com/lyx0/gophercises/blackjack_ai/blackjack"
	"github.com/lyx0/gophercises/deck"
)

func (ai *basicAI) Bet(shuffled bool) int {
	panic("not implemented") // TODO: Implement
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	panic("not implemented") // TODO: Implement
}

func (ai *basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	panic("not implemented") // TODO: Implement
}

func main() {
	opts := blackjack.Options{
		Decks:           3,
		Hands:           1,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
