package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
)

func TestCheck(t *testing.T) {
	initialBoard := `
 BR  BK  BB  BQ  Bk  BB  BK  BR 
 BP  BP  BP  BP  __  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  WP  __  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`
	game, err := game.NewGameWithBoard(initialBoard)
	assert.NoError(t, err)

	service := NewCheckService(game)

	iSquare := game.Board[0][3]

	tSquare := game.Board[4][7]

	movement, err := entity.NewMovement(iSquare, tSquare)
	assert.NoError(t, err)

	_, isCheck := service.HandleCheck(movement)
	assert.True(t, isCheck)
}
