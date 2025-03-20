package entity

import (
	"fmt"

	"github.com/vctaragao/chess/pkg/chess/helper"
)

type Square struct {
	helper.Position
	Piece *Piece
}

func NewSquare(line, column int) *Square {
	return &Square{
		Position: helper.Position{
			Line:   line,
			Column: column,
		},
	}
}

func (s *Square) IsEmpty() bool {
	return s.Piece == nil
}

func (s *Square) SetPiece(p *Piece) {
	s.Piece = p
	if p != nil {
		s.Piece.Square = s
	}
}

func (s *Square) RemovePiece() {
	s.Piece = nil
}

func (s *Square) String() string {
	return fmt.Sprintf("(%d, %d), Piece: %s\n", s.Line, s.Column, s.Piece)

}
