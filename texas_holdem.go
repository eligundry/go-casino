package casino

import (
	"github.com/adamclerk/deck"
)

type TexasHoldem struct {
	deck      *deck.Deck
	players   *Players
	dealer    *Player
	community *deck.Deck
	pot       int
}

func NewTexasHoldemGame(players *Players) TexasHoldem {
	cards, _ := deck.New()
	community, _ := deck.New(deck.Empty)
	game := TexasHoldem{
		players:   players,
		deck:      cards,
		community: community,
		dealer:    players.Dealer(),
	}

	// Deal all the players 2 cards
	DealCardsToPlayers(2, game.deck, game.players)

	return game
}

func (game *TexasHoldem) DealToCommunity(count int) {
	game.deck.Deal(count, game.community)
}
