package service

import (
	"errors"

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

func (s *MovementService) HandleMovement(m entity.Movement) error {
	if !m.IsValid() {
		return errors.New("invalid movement")
	}

	if m.IsCapture() {
		s.HandlePoints(m)
	}

	m.TargetSquare.SetPiece(m.GetPiece())
	s.SetSquare(m.TargetY(), m.TargetX(), m.TargetSquare)

	m.InitialSquare.SetPiece(entity.NewEmptyPiece())
	s.SetSquare(m.InitialY(), m.InitialX(), m.InitialSquare)

	return nil
}
