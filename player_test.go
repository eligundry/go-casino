package casino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftDealer(t *testing.T) {
	players := NewPlayers(2)
	assert.True(t, players.players[0].is_dealer, "The 1st player is the dealer.")

	// Shift the dealer to the second player
	players.ShiftDealer()
	assert.False(t, players.players[0].is_dealer, "The 1st player is not the dealer.")
	assert.True(t, players.players[1].is_dealer, "The 2nd player is the dealer.")

	// Shift the dealer back to the first player
	players.ShiftDealer()
	assert.True(t, players.players[0].is_dealer, "The 1st player is the dealer.")
	assert.False(t, players.players[1].is_dealer, "The 2nd player is not the dealer.")
}
