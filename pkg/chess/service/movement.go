package service

import (
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
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
	m.InitialSquare.RemovePiece()
}
