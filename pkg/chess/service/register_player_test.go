package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/chess/pkg/chess/game"
)

func TestRegisterPlayer(t *testing.T) {
	g := game.NewGame()
	service := NewRegisterPlayerService(g)

	err := service.Execute("player1")
	assert.NoError(t, err)

	assert.Equal(t, g.WPlayer.Nick, "player1")

	err = service.Execute("player2")
	assert.NoError(t, err)

	assert.Equal(t, g.BPlayer.Nick, "player2")

	err = service.Execute("player3")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrMaxPlayer)
}
