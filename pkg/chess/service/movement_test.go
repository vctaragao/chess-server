package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
)

func TestMovement(t *testing.T) {
	g, err := game.NewGame()
	assert.NoError(t, err)

	mService := NewMovementService(g)

	iSquare := g.Board[1][2]
	tSquare := g.Board[3][2]

	piece := iSquare.Piece

	movement, err := entity.NewMovement(iSquare, tSquare)
	assert.NoError(t, err)

	mService.HandleMovement(movement)

	assert.Nil(t, iSquare.Piece)
	assert.Equal(t, tSquare.Piece, piece)
}
