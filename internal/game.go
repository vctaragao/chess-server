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
	g.move(movement)
}

func (g *game) move(m entity.Movement) {

}

func (g *game) handlePawnMovement(m entity.Movement) {

}

func (g *game) GetSquare(y, x int) entity.Square {
	return g.board[y][x]
}

func (g *game) ParseAction(action string) entity.Action {
	return entity.NewActionFromString(action)
}

func (g *game) ParseResult(action string) entity.Result {
	return entity.NewResultFromString(action)

}
