package casino

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTexasHoldemGame(t *testing.T) {
	players := NewPlayers(2)
	game := CreateTexasHoldemGame(players)

	assert.Equal(t, 2, len(game.players), "There are two players!")

	for i := 0; i < 2; i++ {
		assert.Equal(t, 2, players[i].hand.NumberOfCards(), "Each player must have 2 cards.")
	}
}
