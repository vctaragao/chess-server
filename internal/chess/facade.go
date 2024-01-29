package chess

import (
	"github.com/vctaragao/chess-server/internal/chess/entity"
	"github.com/vctaragao/chess-server/internal/chess/game"
	"github.com/vctaragao/chess-server/internal/chess/service"
	"github.com/vctaragao/chess-server/internal/chess/usecases"
)

type facade struct {
	game           *game.Game
	move           *usecases.Move
	registerPlayer *usecases.RegisterPlayer
}

func NewGame() *facade {
	game := game.NewGame()

	checkService := service.NewCheckService(game)
	movementService := service.NewMovementService(game)
	checkMateService := service.NewCheckMateService(game)
	registerPlayerService := service.NewRegisterPlayerService(game)

	facade := &facade{
		game:           game,
		registerPlayer: usecases.NewRegisterPlayer(registerPlayerService),
		move:           usecases.NewMove(game, movementService, checkService, checkMateService),
	}

	return facade
}

func (f *facade) RegisterPlayer(nick string) error {
	return f.registerPlayer.Execute(nick)
}

func (c *facade) Render() {
	c.game.Render()
}

func (f *facade) Move(iSquare, tSquare *entity.Square) (err error) {
	return f.move.Execute(iSquare, tSquare)
}

func (f *facade) GetSquare(y, x int) *entity.Square {
	return f.game.GetSquare(y, x)
}

func (f *facade) GetStatus() game.Status {
	return f.game.Status
}
