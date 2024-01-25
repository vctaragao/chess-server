package chess

import "github.com/vctaragao/chess-server/internal/chess/entity"

type (
	Action int
	Result int

	Movement struct {
		InitialSquare entity.Square
		TargetSquare  entity.Square
		Action        Action
		Result        Result
	}
)

const (
	Move Action = iota
	Capture
)

func NewMovement(initialSquare, targetSquare entity.Square) Movement {
	return Movement{
		InitialSquare: initialSquare,
		TargetSquare:  targetSquare,
	}
}

func NewActionFromString(action string) Action {
	if action == "capture" {
		return Capture
	}

	return Move
}

func (m *Movement) TargetY() int {
	return m.TargetSquare.Y
}

func (m *Movement) TargetX() int {
	return m.TargetSquare.X
}

func (m *Movement) InitialY() int {
	return m.InitialSquare.Y
}

func (m *Movement) InitialX() int {
	return m.InitialSquare.X
}

func (m *Movement) GetPiece() *entity.Piece {
	return m.InitialSquare.Piece
}

func (m *Movement) GetTargetPiece() *entity.Piece {
	return m.TargetSquare.Piece
}

func (m *Movement) IsValid() bool {
	if !m.TargetSquare.IsEmpty() && m.TargetSquare.Piece.Color == m.InitialSquare.Piece.Color {
		return false
	}

	return true
}

func (m *Movement) IsCapture() bool {
	return !m.TargetSquare.IsEmpty() && m.TargetSquare.Piece.Color != m.InitialSquare.Piece.Color
}
