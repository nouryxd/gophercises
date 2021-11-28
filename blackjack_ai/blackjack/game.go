package blackjack

import (
	"fmt"

	"github.com/lyx0/gophercises/deck"
)

const (
	statePlayerTurn State = iota
	stateDealerTurn
	stateHandOver
)

type State int8

type Game struct {
	deck   []deck.Card
	state  state
	player Hand
	dealer Hand
}

type Move func(*Game) Game

func Hit(gs *Game) {
}

func Stand(gs *Game) {
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func Deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func (g *Game) Play(ai AI) {
	g.deck = deck.New(deck.Deck(3), deck.Shuffle)

	for i := 0; i < 10; i++ {
		gs = Deal(gs)

		var input string
		for gs.State == StatePlayerTurn {
			fmt.Println("Player: ", gs.Player)
			fmt.Println("Dealer: ", gs.Dealer.DealerString())
			fmt.Printf("Your cards are worth %d\n", gs.Player.Score())
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			fmt.Printf("\n")
			switch input {
			case "h":
				gs = Hit(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid option:", input)
			}
		}

		// If dealer score <= 16, we hit
		// If dealer has a soft 17, then we hit. (Ace as
		// 11 points and 16 points)
		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Hit(gs)
			} else {
				gs = Stand(gs)
			}
		}

		gs = EndHand(gs)
	}
}
