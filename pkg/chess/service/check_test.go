package service

import (
	"testing"

	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
)

func TestCheck(t *testing.T) {
	boardInitial := `
 BR  BK  BB  __  Bk  BB  BK  BR 
 BP  BP  BP  BP  BP  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  BQ  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  __  WP  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`
	g := game.NewGame()

	service := NewCheckService(g)

	iSquare := g.Board[5][3]
	tSquare := g.Board[4][3]

	piece := iSquare.Piece

	movement := entity.NewMovement(iSquare, tSquare)

	service.HandleCheck(movement)
}
