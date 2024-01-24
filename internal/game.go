package internal

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/vctaragao/chess-server/internal/entity"
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
	board   entity.Board
	wPlayer *entity.Player
	bPlayer *entity.Player
	turn    entity.Color
}

func NewGame() *Game {
	return &Game{
		board: entity.NewBoard(),
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
		g.bPlayer = entity.NewPlayer(nick)
	} else {
		g.wPlayer = entity.NewPlayer(nick)
	}

	return nil
}

func (g *Game) NewMovement(iSquare, tSquare *entity.Square) entity.Movement {
	return entity.NewMovement(*iSquare, *tSquare)
}

func (g *Game) HandleMovement(m entity.Movement) error {
	if !m.IsValid() {
		return errors.New("invalid movement")
	}

	g.move(m)
	g.changeTurn()

	return nil
}

func (g *Game) move(m entity.Movement) {
	if m.IsCapture() {
		g.handlePoints(m)
	}

	m.TargetSquare.SetPiece(m.GetPiece())
	g.setSquare(m.TargetY(), m.TargetX(), &m.TargetSquare)

	m.InitialSquare.SetPiece(entity.NewEmptyPiece())
	g.setSquare(m.InitialY(), m.InitialX(), &m.InitialSquare)

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

func (g *Game) isCheckMate(m entity.Movement) bool {
	// check if king can move or if any piece can block the attack or capture the attacking piece
	if g.canKingMove(m) || g.canPieceBeBlocked(m) || g.canPieceBeCaptured(m) {
		return false
	}

	return true
}

func (g *Game) canPieceBeBlocked(m entity.Movement) bool {
	piece := m.GetPiece()
	possibleSquaresToBeBlocked := g.getAttackingSquares(piece)

	for _, square := range possibleSquaresToBeBlocked {
		if g.canPieceMoveToSquare(m, square) {
			return true
		}
	}

	return false
}

func (g *Game) canPieceMoveToSquare(m entity.Movement, square *entity.Square) bool {
	// check if piece can move to square
	// if it can, check if it can capture the attacking piece
	// if it can, return true

	// if it cant, return false
	return false
}

func (g *Game) canPieceBeCaptured(m entity.Movement) bool {
	return true
}

func (g *Game) canKingMove(m entity.Movement) bool {
	piece := m.GetPiece()

	kColor := entity.White
	if piece.IsWhite() {
		kColor = entity.Black
	}

	kingSquare := g.getKingSquare(kColor)
	kingPossibleSquares := g.getKingAttackingSquares(kingSquare.GetPiece())

	attackingSquares := g.getAllAttackingSquares(piece.GetColor())

	for _, square := range kingPossibleSquares {
		if slices.Contains(attackingSquares, square) {
			return true
		}
	}

	return false
}

func (g *Game) getAllAttackingSquares(color entity.Color) []*entity.Square {
	var squares []*entity.Square

	pieces := g.getAllPieces(color)
	for _, piece := range pieces {
		squares = append(squares, g.getAttackingSquares(piece)...)
	}

	return squares
}

func (g *Game) getAllPieces(color entity.Color) []*entity.Piece {
	var pieces []*entity.Piece

	for _, row := range g.board {
		for _, square := range row {
			if square.IsEmpty() {
				continue
			}

			piece := square.GetPiece()
			if piece.HasColor(color) {
				pieces = append(pieces, piece)
			}
		}
	}

	return pieces
}

func (g *Game) isCheck(m entity.Movement) bool {
	piece := m.GetPiece()

	kColor := entity.White
	if piece.IsWhite() {
		kColor = entity.Black
	}

	kingSquare := g.getKingSquare(kColor)
	attackingSquares := g.getAttackingSquares(piece)

	return slices.Contains(attackingSquares, kingSquare)
}

func (g *Game) getAttackingSquares(piece *entity.Piece) []*entity.Square {
	switch piece.GetType() {
	case entity.Pawn:
		return g.getPawnAttackingSquares(piece)
	case entity.Knight:
		return g.getKnightAttackingSquares(piece)
	case entity.Bishop:
		return g.getBishopAttackingSquares(piece)
	case entity.Rook:
		return g.getRookAttackingSquares(piece)
	case entity.Queen:
		return g.getQueenAttackingSquares(piece)
	case entity.King:
		return g.getKingAttackingSquares(piece)
	default:
		return []*entity.Square{}
	}
}

func (g *Game) getKingAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.GetSquare().Y, piece.GetSquare().X

	return append(squares, g.GetSquare(y-1, x-1), g.GetSquare(y-1, x), g.GetSquare(y-1, x+1))
}

func (g *Game) getRookAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.GetSquare().Y, piece.GetSquare().X

	for down := y - 1; down >= 0; down-- {
		squares = append(squares, g.GetSquare(down, x))

		if !g.GetSquare(down, x).IsEmpty() {
			break
		}
	}

	for up := y + 1; up <= 7; up++ {
		squares = append(squares, g.GetSquare(up, x))
		if !g.GetSquare(up, x).IsEmpty() {
			break
		}
	}

	for left := x - 1; left >= 0; left-- {
		squares = append(squares, g.GetSquare(y, left))
		if !g.GetSquare(y, left).IsEmpty() {
			break
		}
	}

	for right := x + 1; right <= 7; right++ {
		squares = append(squares, g.GetSquare(y, right))
		if !g.GetSquare(y, right).IsEmpty() {
			break
		}
	}

	return squares
}

func (g *Game) getQueenAttackingSquares(piece *entity.Piece) []*entity.Square {
	return append(g.getRookAttackingSquares(piece), g.getBishopAttackingSquares(piece)...)
}

func (g *Game) getPawnAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.GetSquare().Y, piece.GetSquare().X

	if piece.IsWhite() {
		return append(squares, g.GetSquare(y-1, x-1), g.GetSquare(y-1, x+1))
	}

	return append(squares, g.GetSquare(y+1, x-1), g.GetSquare(y+1, x+1))
}

func (g *Game) getKnightAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.GetSquare().Y, piece.GetSquare().X

	return append(squares,
		g.GetSquare(y-2, x-1),
		g.GetSquare(y-2, x+1),
		g.GetSquare(y-1, x-2),
		g.GetSquare(y-1, x+2),
		g.GetSquare(y+1, x-2),
		g.GetSquare(y+1, x+2),
		g.GetSquare(y+2, x-1),
		g.GetSquare(y+2, x+1),
	)
}

func (g *Game) getBishopAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.GetSquare().Y, piece.GetSquare().X

	for down, left := y-1, x-1; down >= 0 && left >= 0; down, left = down-1, left-1 {
		squares = append(squares, g.GetSquare(down, left))

		if !g.GetSquare(down, left).IsEmpty() {
			break
		}
	}

	for down, right := y-1, x+1; down >= 0 && right <= 7; down, right = down-1, right+1 {
		squares = append(squares, g.GetSquare(down, right))

		if !g.GetSquare(down, right).IsEmpty() {
			break
		}
	}

	for up, left := y+1, x-1; up <= 7 && left >= 0; up, left = up+1, left-1 {
		squares = append(squares, g.GetSquare(up, left))

		if !g.GetSquare(up, left).IsEmpty() {
			break
		}
	}

	for up, right := y+1, x+1; up <= 7 && right <= 7; up, right = up+1, right+1 {
		squares = append(squares, g.GetSquare(up, right))

		if !g.GetSquare(up, right).IsEmpty() {
			break
		}
	}

	return squares
}

func (g *Game) getKingSquare(color entity.Color) *entity.Square {
	for _, row := range g.board {
		for _, square := range row {
			if square.IsEmpty() {
				continue
			}

			p := square.GetPiece()
			if p.Is(entity.King) && p.HasColor(color) {
				return square
			}
		}
	}

	return nil
}

func (g *Game) handlePoints(m entity.Movement) {
	wOp, bOp := Sub, Add
	tPiece := m.GetTargetPiece()

	if tPiece.IsBlack() {
		wOp, bOp = Add, Sub
	}

	g.changeWhitePoints(tPiece.GetValue(), wOp)
	g.changeBlackPoints(tPiece.GetValue(), bOp)

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

func (g *Game) ParseAction(action string) entity.Action {
	return entity.NewActionFromString(action)
}

func (g *Game) ParseResult(action string) entity.Result {
	return entity.NewResultFromString(action)
}

func (g *Game) changeTurn() {
	if g.turn == entity.White {
		g.turn = entity.Black
		return
	}

	g.turn = entity.White
}
