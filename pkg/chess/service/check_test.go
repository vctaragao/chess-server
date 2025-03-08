package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
)

func TestCheck(t *testing.T) {
	_ = `
 BR  BK  BB  __  Bk  BB  BK  BR 
 BP  BP  BP  BP  BP  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  BQ  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  __  WP  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`
	game, err := game.NewGame()
	assert.NoError(t, err)

	service := NewCheckService(game)

	iSquare := game.Board[5][3]
	tSquare := game.Board[4][3]

	movement := entity.NewMovement(iSquare, tSquare)

	service.HandleCheck(movement)
}
