package entity

type Piece struct {
	color     Color
	isNull    bool
	square    *Square
	pieceType PieceType
}

func NewPiece(c Color, t PieceType) Piece {
	return Piece{
		color:     c,
		pieceType: t,
	}
}

func NewEmptyPiece() Piece {
	return Piece{
		isNull:    true,
		pieceType: None,
	}
}

func (p Piece) GetType() PieceType {
	return p.pieceType
}

func (p *Piece) SetSquare(square *Square) {
	p.square = square
}
