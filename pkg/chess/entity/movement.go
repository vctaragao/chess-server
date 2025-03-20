package entity

import (
	"errors"
	"fmt"
)

type (
	Action int
	Result int

	Movement struct {
		InitialSquare *Square
		TargetSquare  *Square
		Action        Action
		Result        Result
	}
)

var ErrInvalidMovement = errors.New("invalid movement")

const (
	Move Action = iota
	Capture
)

func NewMovement(initialSquare, targetSquare *Square) (*Movement, error) {
	m := &Movement{
		InitialSquare: initialSquare,
		TargetSquare:  targetSquare,
	}

	if !m.IsValid() {
		return m, ErrInvalidMovement
	}

	return m, nil
}

func NewActionFromString(action string) Action {
	if action == "capture" {
		return Capture
	}

	return Move
}

func (m *Movement) TargetLine() int {
	return m.TargetSquare.Line
}

func (m *Movement) TargetColumn() int {
	return m.TargetSquare.Column
}

func (m *Movement) InitialLine() int {
	return m.InitialSquare.Line
}

func (m *Movement) InitialColumn() int {
	return m.InitialSquare.Column
}

func (m *Movement) GetPiece() *Piece {
	return m.InitialSquare.Piece
}

func (m *Movement) GetTargetPiece() *Piece {
	return m.TargetSquare.Piece
}

// TODO: make this validation function more robust
// to handle other possible cases
func (m *Movement) IsValid() bool {
	if !m.TargetSquare.IsEmpty() && m.TargetSquare.Piece.Color == m.InitialSquare.Piece.Color {
		return false
	}

	return true
}

func (m *Movement) IsCapture() bool {
	return !m.TargetSquare.IsEmpty() && m.TargetSquare.Piece.Color != m.InitialSquare.Piece.Color
}

func (m *Movement) String() string {
	return fmt.Sprintf("iSquare: %s\ntSquare: %s\n", m.InitialSquare, m.TargetSquare)
}
