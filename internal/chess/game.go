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
	if g.isCheck(m) {
		// check if its a checkmate
		// if its a checkmate, set status to checkmate

		if g.isCheckMate(m) {
			g.status = CheckMate
			return
		}

		g.status = Check
	}
}

func (g *Game) isCheckMate(m Movement) bool {
	// check if king can move or if any piece can block the attack or capture the attacking piece
	if g.canKingMove(m) || g.canPieceBeBlocked(m) || g.canPieceBeCaptured(m) {
		return false
	}

	return true
}

func (g *Game) canPieceBeBlocked(m Movement) bool {
	piece := m.GetPiece()

	if piece.Is(entity.Knight) || piece.Is(entity.Pawn) {
		return false
	}

	// check if any piece can move to square that is between the attacking piece and the king
	// get the square that the attacking piece is checking

	piece = m.GetPiece()

	return false
}

func (g *Game) canPieceMoveToSquare(m Movement, square *entity.Square) bool {
	// TODO: implement
	return true
}

func (g *Game) canPieceBeCaptured(m Movement) bool {
	// TODO: implement
	return true
}

func (g *Game) canKingMove(m Movement) bool {
	piece := m.GetPiece()

	kColor := helper.White
	if piece.IsWhite() {
		kColor = helper.Black
	}

	kingSquare := g.getKingSquare(kColor)
	kingPossibleSquares := kingSquare.Piece.AttackingSquares

	attackingSquares := g.getAllAttackingSquares(piece.Color)

	canKingMove := false

	for _, square := range kingPossibleSquares {
		// if square is not empty and has a piece of the same color or it has a piece of a different color that it protected by another piece
		if !square.IsEmpty() && square.Piece.HasColor(kColor) {
			continue
		}
		if slices.Contains(attackingSquares, square) {
			return true
		}
	}

	return false
}

func (g *Game) getAllAttackingSquares(color helper.Color) []*entity.Square {
	var squares []*entity.Square

	pieces := g.getAllPieces(color)
	for _, piece := range pieces {
		squares = append(squares, piece.AttackingSquares...)
	}

	return squares
}

func (g *Game) getAllPieces(color helper.Color) []*entity.Piece {
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

func (g *Game) isCheck(m Movement) bool {
	piece := m.GetPiece()

	kColor := helper.White
	if piece.IsWhite() {
		kColor = helper.Black
	}

	kingSquare := g.getKingSquare(kColor)
	return slices.Contains(piece.AttackingSquares, kingSquare)
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
