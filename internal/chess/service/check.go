package service

import (
	"slices"

	"github.com/vctaragao/chess-server/internal/chess/entity"
	"github.com/vctaragao/chess-server/internal/chess/game"
	"github.com/vctaragao/chess-server/internal/chess/helper"
)

type CheckService struct {
	*game.Game
}

func NewCheckService(g *game.Game) *CheckService {
	return &CheckService{
		Game: g,
	}
}

func (s *CheckService) HandleCheck(m *entity.Movement) (*entity.Square, bool) {
	piece := m.GetPiece()

	kColor := helper.White
	if piece.IsWhite() {
		kColor = helper.Black
	}

	// check if its a direct check
	kingSquare := s.getKingSquare(kColor)
	if slices.Contains(piece.AttackingSquares, kingSquare) {
		return kingSquare, true
	}

	// check if its a discovered check
	pieces := s.GetAllPiecesByColor(piece.Color)
	for _, p := range pieces {
		if p.Is(entity.King) || p.Is(entity.Pawn) || p.Is(entity.Knight) {
			continue
		}

		if p.Is(entity.Bishop) || p.Is(entity.Rook) || p.Is(entity.Queen) {
			if slices.Contains(p.AttackingSquares, kingSquare) {
				s.Status = game.Check
				return kingSquare, true
			}
		}
	}

	return kingSquare, false
}

func (s *CheckService) getKingSquare(color helper.Color) *entity.Square {
	for _, row := range s.Board {
		for _, square := range row {
			if square.IsEmpty() {
				continue
			}

			p := square.Piece
			if p.Is(entity.King) && p.HasColor(color) {
				return square
			}
		}
	}

	return nil
}
