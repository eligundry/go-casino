package casino

import (
	"github.com/adamclerk/deck"
)

type Player struct {
	hand        deck.Deck
	interactive bool
	money       int
}

func NewPlayer() Player {
	hand, _ := deck.New(deck.Empty)
	player := Player{
		interactive: false,
		money:       0,
		hand:        *hand,
	}

	return player
}

func NewPlayers(count int) []Player {
	var players []Player

	for i := 0; i < count; i++ {
		players = append(players, NewPlayer())
	}

	return players
}

func DealCardsToPlayers(count int, cards *deck.Deck, players []Player) {
	// var hands []deck.Deck

	// for i := 0; i < len(players); i++ {
	// 	cards.Deal(count, players[i].hand)
	// }
}
