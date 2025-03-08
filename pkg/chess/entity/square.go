package entity

import (
	"github.com/vctaragao/chess/pkg/chess/helper"
)

type Square struct {
	helper.Position
	Piece *Piece
}

func NewSquare(y, x int) *Square {
	return &Square{
		Position: helper.Position{
			X: x,
			Y: y,
		},
	}
}

func (s *Square) IsEmpty() bool {
	return s.Piece == nil
}

func (s *Square) SetPiece(p *Piece) {
	s.Piece = p
}

func (s *Square) RemovePiece() {
	s.Piece = nil
}
