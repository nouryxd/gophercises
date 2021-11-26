package main

import (
	"fmt"
	"strings"

	deck "github.com/lyx0/gophercises/deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = cards[0], cards[1:]
			*hand = append(*hand, card)
		}
	}
	fmt.Println("Player: ", player)
	fmt.Println("Dealer: ", dealer)
}
