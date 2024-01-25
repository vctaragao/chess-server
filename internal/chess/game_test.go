package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterPlayer(t *testing.T) {
	game := NewGame()

	err := game.RegisterPlayer("victor")
	assert.NoError(t, err)
	assert.NotNil(t, game.wPlayer)

	err = game.RegisterPlayer("test")
	assert.NoError(t, err)
	assert.NotNil(t, game.bPlayer)
}

func TestRegisterPlayerError(t *testing.T) {
	game := NewGame()

	err := game.RegisterPlayer("victor")
	assert.NoError(t, err)
	assert.NotNil(t, game.wPlayer)

	err = game.RegisterPlayer("pedro")
	assert.NoError(t, err)
	assert.NotNil(t, game.bPlayer)

	err = game.RegisterPlayer("victor")
	assert.EqualError(t, err, ErrMaxPlayer.Error())
}
