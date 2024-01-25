package entity

type PieceType int

const (
	None PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p PieceType) String() string {
	return []string{"__", "P", "K", "B", "R", "Q", "k"}[p]
}

func PieceTypeFromString(p string) PieceType {
	switch p {
	case "P":
		return Pawn
	case "K":
		return Knight
	case "B":
		return Bishop
	case "R":
		return Rook
	case "Q":
		return Queen
	case "k":
		return King
	default:
		return None
	}
}
