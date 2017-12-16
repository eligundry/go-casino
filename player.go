package casino

import (
	"github.com/adamclerk/deck"
)

type Player struct {
	hand        *deck.Deck
	interactive bool
	money       int
	active      bool
	is_dealer   bool
}

type Players struct {
	players []Player
}

func (p *Players) Len() int {
	return len(p.players)
}

func (p *Players) Dealer() *Player {
	var target Player

	for i := 0; i < p.Len(); i++ {
		if p.players[i].is_dealer {
			target = p.players[i]
			break
		}
	}

	return &target
}

func (p *Players) ShiftDealer() *Player {
	var dealer *Player
	foundDealer := false

	for i := 0; i < p.Len(); i++ {
		foundDealer = p.players[i].is_dealer

		if foundDealer {
			if i+1 == p.Len() {
				dealer = &p.players[0]
			} else {
				dealer = &p.players[i+1]
			}

			p.players[i].is_dealer = false
			dealer.is_dealer = true
			break
		}
	}

	return dealer
}

func NewPlayer() Player {
	hand, _ := deck.New(deck.Empty)
	player := Player{
		interactive: false,
		money:       100,
		hand:        hand,
		active:      true,
		is_dealer:   false,
	}

	return player
}

func NewPlayers(count int) Players {
	players := Players{}

	for i := 0; i < count; i++ {
		player := NewPlayer()

		if i == 0 {
			player.is_dealer = true
		}

		players.players = append(players.players, player)
	}

	return players
}

func DealCardsToPlayers(count int, cards *deck.Deck, players *Players) {
	for i := 0; i < players.Len(); i++ {
		cards.Deal(count, players.players[i].hand)
	}
}

func (player *Player) Bet(amount int, game *TexasHoldem) bool {
	if player.money < amount {
		return false
	}

	player.money -= amount
	game.pot += amount

	return true
}

func (player *Player) Win(pot int) {
	player.money += pot
}

func (player *Player) Lose() bool {
	return true
}
