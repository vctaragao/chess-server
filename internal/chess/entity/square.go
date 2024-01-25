package entity

import (
	"github.com/vctaragao/chess-server/internal/chess/helper"
)

type Square struct {
	helper.Position
	Empty bool
	Piece *Piece
}

func NewSquare(y, x int) *Square {
	return &Square{
		Empty: true,
		Position: helper.Position{
			X: x,
			Y: y,
		},
	}
}

func (s *Square) IsEmpty() bool {
	return s.Empty
}

func (s *Square) SetPiece(p *Piece) {
	if !p.IsNull {
		s.Empty = false
	}

	s.Piece = p
}
