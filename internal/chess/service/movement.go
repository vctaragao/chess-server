package service

import (
	"github.com/vctaragao/chess-server/internal/chess/entity"
	"github.com/vctaragao/chess-server/internal/chess/game"
)

type MovementService struct {
	*game.Game
}

func NewMovementService(g *game.Game) *MovementService {
	return &MovementService{
		Game: g,
	}
}

func (s *MovementService) HandleMovement(m *entity.Movement) {
	m.TargetSquare.SetPiece(m.InitialSquare.Piece)
	m.InitialSquare.SetPiece(entity.NewEmptyPiece())
}
