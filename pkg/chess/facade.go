package chess

import (
	"log"

	"github.com/vctaragao/chess/pkg/chess/board"
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
	"github.com/vctaragao/chess/pkg/chess/service"
	"github.com/vctaragao/chess/pkg/chess/usecases"
)

type GameState struct {
	Board [][]string
}

type facade struct {
	game           *game.Game
	move           *usecases.Move
	registerPlayer *usecases.RegisterPlayer
}

func NewGame() (*facade, error) {
	game, err := game.NewGame()
	if err != nil {
		return nil, err
	}

	checkService := service.NewCheckService(game)
	movementService := service.NewMovementService(game)
	checkMateService := service.NewCheckMateService(game)
	registerPlayerService := service.NewRegisterPlayerService(game)

	facade := &facade{
		game:           game,
		registerPlayer: usecases.NewRegisterPlayer(registerPlayerService),
		move:           usecases.NewMove(game, movementService, checkService, checkMateService),
	}

	return facade, nil
}

func (f *facade) RegisterPlayer(nick string) error {
	return f.registerPlayer.Execute(nick)
}

func (c *facade) Render() {
	c.game.Render()
}

func (c *facade) Board() board.Board {
	return c.game.Board
}

func (f *facade) Move(iLine, iColumn, tLine, tColumn int) (err error) {
	return f.move.Execute(iLine, iColumn, tLine, tColumn)
}

func (f *facade) GetSquare(line, column int) *entity.Square {
	return f.game.GetSquare(line, column)
}

func (f *facade) GetStatus() game.Status {
	return f.game.Status
}

func (f *facade) GetState() GameState {
	log.Println("GetState")
	return GameState{
		Board: f.game.Board.State(),
	}
}
