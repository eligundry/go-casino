package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTexasHoldemGame(t *testing.T) {
	players := NewPlayers(2)
	game := NewTexasHoldemGame(&players)

	assert.Equal(t, 2, game.players.Len(), "There are two players.")
	assert.Equal(t, 48, game.deck.NumberOfCards(), "The deck has 52 cards, but 2 were dealt to each player.")

	for i := 0; i < 2; i++ {
		assert.Equal(t, 2, players.players[i].hand.NumberOfCards(), "Each player must have 2 cards.")
	}
}
