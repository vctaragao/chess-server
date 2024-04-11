package entity

import (
	"github.com/vctaragao/chess/pkg/chess/helper"
)

type Piece struct {
	Value            int
	IsNull           bool
	Square           *Square
	AttackingSquares []*Square
	PieceType        PieceType
	Color            helper.Color
	ProtecedBy       *Piece
}

func NewPiece(c helper.Color, t PieceType) *Piece {
	return &Piece{
		Color:     c,
		PieceType: t,
	}
}

func NewEmptyPiece() *Piece {
	return &Piece{
		IsNull:    true,
		PieceType: None,
	}
}

func (p *Piece) HasColor(c helper.Color) bool {
	return p.Color == c
}

func (p *Piece) IsWhite() bool {
	return p.Color == helper.White
}

func (p *Piece) IsBlack() bool {
	return p.Color == helper.Black
}

func (p *Piece) Is(t PieceType) bool {
	return p.PieceType == t
}

func (p *Piece) IsProteced() bool {
	return p.ProtecedBy != nil
}
