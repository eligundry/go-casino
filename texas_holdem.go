package casino

import (
	"github.com/adamclerk/deck"
)

type TexasHoldem struct {
	deck      deck.Deck
	players   []Player
	community []deck.Card
	pot       int
}

func CreateTexasHoldemGame(players []Player) TexasHoldem {
	cards, _ := deck.New()
	game := TexasHoldem{
		players: players,
		deck:    *cards,
	}

	// Deal all the players 2 cards
	for i := 0; i < len(players); i++ {
		// game.deck.Deal(2, players[i].hand)
	}

	return game
}
