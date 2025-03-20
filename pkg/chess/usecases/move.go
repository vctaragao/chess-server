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

func (m *Move) Execute(iLine, iColumn, tLine, tColumn int) (err error) {
	iSquare := m.Game.Square(iLine, iColumn)
	tSquare := m.Game.Square(tLine, tColumn)

	movement, err := entity.NewMovement(iSquare, tSquare)
	if err != nil {
		return err
	}

	m.MovementService.HandleMovement(movement)

	m.Game.Board.UpdateAttackingSquares()

	kSquare, isCheck := m.CheckService.HandleCheck(movement)
	if isCheck {
		m.CheckMateService.HandleCheckMate(movement, kSquare)
	}

	m.Game.ChangeTurn()

	return nil
}
