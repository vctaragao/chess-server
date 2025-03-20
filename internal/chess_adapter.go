package internal

import (
	"github.com/vctaragao/chess/internal/entity"
)

const (
	Empty   = "__"
	WPawn   = "WP"
	WRook   = "WR"
	WKnight = "WK"
	WBishop = "WB"
	WQueen  = "WQ"
	WKing   = "Wk"

	BPawn   = "BP"
	BRook   = "BR"
	BKnight = "BK"
	BBishop = "BB"
	BQueen  = "BQ"
	BKing   = "Bk"
)

type ChessAdapter struct{}

func parsePiece(p string) entity.Piece {
	switch p {
	case WPawn:
		return entity.WPawn
	case WRook:
		return entity.WRook
	case WKnight:
		return entity.WKnight
	case WBishop:
		return entity.WBishop
	case WQueen:
		return entity.WQueen
	case WKing:
		return entity.WKing
	case BPawn:
		return entity.BPawn
	case BRook:
		return entity.BRook
	case BKnight:
		return entity.BKnight
	case BBishop:
		return entity.BBishop
	case BQueen:
		return entity.BQueen
	case BKing:
		return entity.BKing
	default:
		return entity.None
	}
}
