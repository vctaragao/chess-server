package chess

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/vctaragao/chess-server/internal/chess/board"
	"github.com/vctaragao/chess-server/internal/chess/entity"
	"github.com/vctaragao/chess-server/internal/chess/helper"
)

var ErrMaxPlayer = errors.New("game already full of players")

type Operation int

const (
	Add Operation = iota
	Sub
)

type Status int

const (
	None Status = iota
	Check
	CheckMate
)

type Game struct {
	status  Status
	wPlayer *Player
	bPlayer *Player
	board   board.Board
	turn    helper.Color
}

func NewGame() *Game {
	return &Game{
		board: board.NewBoard(),
	}
}

func (g *Game) Render() {
	fmt.Printf("%s: %d\n", g.bPlayer.Nick, g.bPlayer.Points)
	fmt.Print(strings.TrimPrefix(g.board.Show(), "\n"))
	fmt.Printf("%s: %d\n\n", g.wPlayer.Nick, g.wPlayer.Points)

	fmt.Println("Player turn: ", g.turn)
}

func (g *Game) RegisterPlayer(nick string) error {
	if g.wPlayer != nil && g.bPlayer != nil {
		return ErrMaxPlayer
	}

	if g.wPlayer != nil {
		g.bPlayer = NewPlayer(nick)
	} else {
		g.wPlayer = NewPlayer(nick)
	}

	return nil
}

func (g *Game) NewMovement(iSquare, tSquare *entity.Square) Movement {
	return NewMovement(*iSquare, *tSquare)
}

func (g *Game) HandleMovement(m Movement) error {
	if !m.IsValid() {
		return errors.New("invalid movement")
	}

	g.move(m)
	g.changeTurn()

	return nil
}

func (g *Game) move(m Movement) {
	if m.IsCapture() {
		g.handlePoints(m)
	}

	m.TargetSquare.SetPiece(m.GetPiece())
	g.setSquare(m.TargetY(), m.TargetX(), &m.TargetSquare)

	m.InitialSquare.SetPiece(entity.NewEmptyPiece())
	g.setSquare(m.InitialY(), m.InitialX(), &m.InitialSquare)

	g.board.UpdateAttackingSquares()

	// check if its a check
	kSquare, isCheck := g.isCheck(m)
	if isCheck {
		// check if its a checkmate
		// if its a checkmate, set status to checkmate

		cms := NewCheckMateService(g, m, kSquare)

		if cms.IsCheckMate() {
			g.status = CheckMate
			return
		}

		g.status = Check
	}
}

func (g *Game) getAllAttackingSquares(color helper.Color) []*entity.Square {
	var squares []*entity.Square

	pieces := g.getAllPiecesByColor(color)
	for _, piece := range pieces {
		squares = append(squares, piece.AttackingSquares...)
	}

	return squares
}

func (g *Game) getAllPiecesByColor(color helper.Color) []*entity.Piece {
	var pieces []*entity.Piece

	for _, row := range g.board {
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

func (g *Game) isCheck(m Movement) (*entity.Square, bool) {
	piece := m.GetPiece()

	kColor := helper.White
	if piece.IsWhite() {
		kColor = helper.Black
	}

	kingSquare := g.getKingSquare(kColor)
	return kingSquare, slices.Contains(piece.AttackingSquares, kingSquare)
}

func (g *Game) getKingSquare(color helper.Color) *entity.Square {
	for _, row := range g.board {
		for _, square := range row {
			if square.IsEmpty() {
				continue
			}

			p := square.Piece
			if p.Is(entity.King) && p.HasColor(color) {
				return square
			}
		}
	}

	return nil
}

func (g *Game) handlePoints(m Movement) {
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
		g.wPlayer.Points += points
		return
	}

	g.wPlayer.Points -= points
}

func (g *Game) changeBlackPoints(points int, op Operation) {
	if op == Add {
		g.bPlayer.Points += points
	}

	g.bPlayer.Points -= points
}

func (g *Game) GetSquare(y, x int) *entity.Square {
	return g.board[y][x]
}

func (g *Game) setSquare(y, x int, square *entity.Square) {
	g.board[y][x] = square
}

func (g *Game) ParseAction(action string) Action {
	return NewActionFromString(action)
}

func (g *Game) changeTurn() {
	if g.turn == helper.White {
		g.turn = helper.Black
		return
	}

	g.turn = helper.White
}
