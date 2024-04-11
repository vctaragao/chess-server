package usecases

import (
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
	"github.com/vctaragao/chess/pkg/chess/service"
)

type Move struct {
	Game             *game.Game
	MovementService  *service.MovementService
	CheckService     *service.CheckService
	CheckMateService *service.CheckMateService
}

func NewMove(g *game.Game, m *service.MovementService, c *service.CheckService, cm *service.CheckMateService) *Move {
	return &Move{
		Game:             g,
		MovementService:  m,
		CheckService:     c,
		CheckMateService: cm,
	}
}

func (m *Move) Execute(iSquare, tSquare *entity.Square) (err error) {
	movement := entity.NewMovement(iSquare, tSquare)
	m.MovementService.HandleMovement(movement)

	m.Game.Board.UpdateAttackingSquares()
	m.Game.ChangeTurn()

	kSquare, isCheck := m.CheckService.HandleCheck(movement)
	if !isCheck {
		return nil
	}

	m.CheckMateService.HandleCheckMate(movement, kSquare)

	return nil
}
