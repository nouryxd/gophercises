package blackjack

import (
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

func Hit(g *Game) {
}

func Stand(g *Game) {
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func deal(g *Game) {
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
		deal(g)

		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}

		// If dealer score <= 16, we hit
		// If dealer has a soft 17, then we hit. (Ace as
		// 11 points and 16 points)
		for g.state == stateDealerTurn {
			if g.dealer.Score() <= 16 || (g.dealer.Score() == 17 && g.dealer.MinScore() != 17) {
				g = Hit(g)
			} else {
				g = Stand(g)
			}
		}

		g = EndHand(g)
	}
}

func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// Ace is currently worth 1 point, here we
			// change it so that it's worth 11
			return minScore + 10
		}
	}
	return minScore
}

func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
	if minScore != score {
		return true
	}
	return false
}

func minScore(hand ...deck.Card) int {
	score := 0

	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
