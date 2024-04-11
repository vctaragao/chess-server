package internal

import (
	"log"

	"github.com/vctaragao/chess/internal/entity"
	"github.com/vctaragao/chess/pkg/chess"
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

type ChessAdapter struct {
}

func (a *ChessAdapter) NewGame() [][]entity.Piece {
	log.Println("NewGame")
	game := chess.NewGame()
	log.Println("game", game)

	gameState := game.GetState()

	state := make([][]entity.Piece, 8)
	for i := range state {
		state[i] = make([]entity.Piece, 8)
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			log.Println("Parsing piece", gameState.Board[y][x])
			state[y][x] = parsePiece(gameState.Board[y][x])
		}
	}

	return state
}

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
