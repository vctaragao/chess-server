package game

import (
	"fmt"
	"strings"

	"github.com/vctaragao/chess/pkg/chess/board"
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/helper"
)

type Operation int

const (
	Add Operation = iota
	Sub
)

type Game struct {
	Status  Status
	WPlayer *entity.Player
	BPlayer *entity.Player
	Board   board.Board
	Turn    helper.Color
}

func NewGame() (*Game, error) {
	game := &Game{}

	board, err := board.NewBoard()
	if err != nil {
		return game, err
	}

	game.Board = board

	return game, nil
}

func NewGameWithBoard(boardStr string) (*Game, error) {
	game := &Game{}

	board, err := board.NewBoardFromString(boardStr)
	if err != nil {
		return game, err
	}

	game.Board = board

	return game, nil
}

func (g *Game) Render() {
	fmt.Printf("%s: %d\n", g.BPlayer.Nick, g.BPlayer.Points)
	fmt.Print(strings.TrimPrefix(g.Board.Show(), "\n"))
	fmt.Printf("%s: %d\n\n", g.WPlayer.Nick, g.WPlayer.Points)

	fmt.Println("Player turn: ", g.Turn)
}

func (g *Game) GetAllAttackingSquares(color helper.Color) []*entity.Square {
	var squares []*entity.Square

	pieces := g.GetAllPiecesByColor(color)
	for _, piece := range pieces {
		squares = append(squares, piece.AttackingSquares...)
	}

	return squares
}

func (g *Game) GetAllPiecesByColor(color helper.Color) []*entity.Piece {
	var pieces []*entity.Piece

	for _, row := range g.Board {
		for _, square := range row {
			if square.IsEmpty() {
				continue
			}

			piece := square.Piece
			if piece.HasColor(color) {
				pieces = append(pieces, piece)
			}
		}
	}

	return pieces
}

func (g *Game) HandlePoints(m *entity.Movement) {
	wOp, bOp := Sub, Add
	tPiece := m.GetTargetPiece()

	if tPiece.IsBlack() {
		wOp, bOp = Add, Sub
	}

	g.changeWhitePoints(tPiece.Value, wOp)
	g.changeBlackPoints(tPiece.Value, bOp)
}

func (g *Game) changeWhitePoints(points int, op Operation) {
	if op == Add {
		g.WPlayer.Points += points
		return
	}

	g.WPlayer.Points -= points
}

func (g *Game) changeBlackPoints(points int, op Operation) {
	if op == Add {
		g.BPlayer.Points += points
	}

	g.BPlayer.Points -= points
}

func (g *Game) GetSquare(y, x int) *entity.Square {
	return g.Board[y][x]
}

func (g *Game) SetSquare(y, x int, square *entity.Square) {
	g.Board[y][x] = square
}

func (g *Game) ParseAction(action string) entity.Action {
	return entity.NewActionFromString(action)
}

func (g *Game) ChangeTurn() {
	if g.Turn == helper.White {
		g.Turn = helper.Black
		return
	}

	g.Turn = helper.White
}
