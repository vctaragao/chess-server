package entity

type Piece int

const (
	None Piece = iota
	WPawn
	WRook
	WKnight
	WBishop
	WQueen
	WKing

	BPawn
	BRook
	BKnight
	BBishop
	BQueen
	BKing
)

func (p Piece) String() string {
	return [...]string{" ", "♙", "♖", "♘", "♗", "♕", "♔", "♟", "♜", "♞", "♝", "♛", "♚"}[p]
}

func ParsePiece(piece string) Piece {
	switch piece {
	case "♙":
		return WPawn
	case "♖":
		return WRook
	case "♘":
		return WKnight
	case "♗":
		return WBishop
	case "♕":
		return WQueen
	case "♔":
		return WKing
	case "♟":
		return BPawn
	case "♜":
		return BRook
	case "♞":
		return BKnight
	case "♝":
		return BBishop
	case "♛":
		return BQueen
	case "♚":
		return BKing
	default:
		return None
	}
}
