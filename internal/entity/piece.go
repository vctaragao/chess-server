package entity

type Piece struct {
	color     Color
	isNull    bool
	square    *Square
	pieceType PieceType
	value     int
}

func NewPiece(c Color, t PieceType) *Piece {
	return &Piece{
		color:     c,
		pieceType: t,
	}
}

func NewEmptyPiece() *Piece {
	return &Piece{
		isNull:    true,
		pieceType: None,
	}
}

func (p Piece) GetType() PieceType {
	return p.pieceType
}

func (p Piece) HasColor(c Color) bool {
	return p.color == c
}

func (p Piece) IsWhite() bool {
	return p.color == White
}

func (p Piece) IsBlack() bool {
	return p.color == Black
}

func (p *Piece) SetSquare(square *Square) {
	p.square = square
}

func (p *Piece) GetValue() int {
	return p.value
}

func (p *Piece) GetSquare() *Square {
	return p.square
}

func (p Piece) Is(t PieceType) bool {
	return p.GetType() == t
}
