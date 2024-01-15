package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vctaragao/chess-server/internal/entity"
)

var ErrMaxPlayer = errors.New("game already full of players")

type game struct {
	board   entity.Board
	wPlayer *entity.Player
	bPlayer *entity.Player
}

func NewGame() game {
	return game{
		board: entity.NewBoard(),
	}
}

func (g *game) Render() {
	fmt.Println(strings.TrimPrefix(g.board.Show(), "\n"))
}

func (g *game) RegisterPlayer() error {
	if g.wPlayer != nil && g.bPlayer != nil {
		return ErrMaxPlayer
	}

	if g.wPlayer != nil {
		g.bPlayer = entity.NewPlayer()
	} else {
		g.wPlayer = entity.NewPlayer()
	}

	return nil
}

func (g *game) HandleMovement(iSquare, tSquare entity.Square, action entity.Action, result entity.Result) {
	movement := entity.NewMovement(iSquare, tSquare, action, result)

	if movement.Action == entity.Move {
		g.handleMove(movement)
	}
}

func (g *game) handleMove(m entity.Movement) {
	tSquare := g.GetSquare(m.TargetY(), m.TargetX())
	tSquare.SetPiece(m.GetPiece())

	g.setSquare(m.TargetY(), m.TargetX(), tSquare)

	iSquare := g.GetSquare(m.InitialY(), m.InitialX())

	emptyPiece := entity.NewEmptyPiece()
	iSquare.SetPiece(&emptyPiece)

	g.setSquare(m.InitialY(), m.InitialX(), iSquare)
}

func (g *game) GetSquare(y, x int) entity.Square {
	return g.board[y][x]
}

func (g *game) setSquare(y, x int, square entity.Square) {
	g.board[y][x] = square
}

func (g *game) ParseAction(action string) entity.Action {
	return entity.NewActionFromString(action)
}

func (g *game) ParseResult(action string) entity.Result {
	return entity.NewResultFromString(action)

}
