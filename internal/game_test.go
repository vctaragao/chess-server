package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterPlayer(t *testing.T) {
	game := NewGame()

	err := game.RegisterPlayer()
	assert.NoError(t, err)
	assert.NotNil(t, game.wPlayer)

	err = game.RegisterPlayer()
	assert.NoError(t, err)
	assert.NotNil(t, game.bPlayer)
}

func TestRegisterPlayerError(t *testing.T) {
	game := NewGame()

	err := game.RegisterPlayer()
	assert.NoError(t, err)
	assert.NotNil(t, game.wPlayer)

	err = game.RegisterPlayer()
	assert.NoError(t, err)
	assert.NotNil(t, game.bPlayer)

	err = game.RegisterPlayer()
	assert.EqualError(t, err, ErrMaxPlayer.Error())
}
