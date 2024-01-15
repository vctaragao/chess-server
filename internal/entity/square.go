package entity

type Square struct {
	Position
	empty bool
	piece *Piece
}

func NewSquare(y, x int) Square {
	return Square{
		empty: true,
		Position: Position{
			x: x,
			y: y,
		},
	}
}

func (s *Square) SetPiece(piece *Piece) {
	s.empty = false
	if piece.isNull {
		s.empty = true
	}

	s.piece = piece
}
