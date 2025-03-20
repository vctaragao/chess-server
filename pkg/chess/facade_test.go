package chess_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/chess/pkg/chess"
)

func TestNewGameBoard(t *testing.T) {
	expectedInitialBoard := `
 BR  BK  BB  BQ  Bk  BB  BK  BR 
 BP  BP  BP  BP  BP  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  WP  WP  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`
	game, err := chess.NewGame()
	assert.NoError(t, err)

	assert.Equal(t, expectedInitialBoard, game.Board().Show())
}

func TestBoardAfterMovement(t *testing.T) {
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
	game, err := chess.NewGameFromStr(initialBoard)
	assert.NoError(t, err)

	err = game.Move(0, 3, 4, 7)
	assert.NoError(t, err)

	expectedFinalBoard := `
 BR  BK  BB  __  Bk  BB  BK  BR 
 BP  BP  BP  BP  __  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  BQ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  WP  __  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`

	assert.Equal(t, expectedFinalBoard, game.Board().Show())

	assert.Equal(t, "Check", game.Status().String())
}
