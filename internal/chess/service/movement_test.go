package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/chess-server/internal/chess/entity"
	"github.com/vctaragao/chess-server/internal/chess/game"
)

func TestMovement(t *testing.T) {
	g := game.NewGame()
	mService := NewMovementService(g)

	iSquare := g.Board[1][2]
	tSquare := g.Board[3][2]

	piece := iSquare.Piece

	movement := entity.NewMovement(iSquare, tSquare)

	mService.HandleMovement(movement)

	assert.Equal(t, tSquare.Piece, piece)
	assert.Equal(t, iSquare.Piece, entity.NewEmptyPiece())
}
