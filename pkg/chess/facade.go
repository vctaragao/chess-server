package chess

import (
	"github.com/vctaragao/chess/pkg/chess/board"
	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
	"github.com/vctaragao/chess/pkg/chess/service"
	"github.com/vctaragao/chess/pkg/chess/usecases"
)

type GameState struct {
	Board [][]string
}

type gameFacade struct {
	game           *game.Game
	move           *usecases.Move
	registerPlayer *usecases.RegisterPlayer
}

func NewGame() (*gameFacade, error) {
	game, err := game.NewGame()
	if err != nil {
		return nil, err
	}

	return newFacade(game), nil
}

func NewGameFromStr(boardStr string) (*gameFacade, error) {
	game, err := game.NewGameWithBoard(boardStr)
	if err != nil {
		return nil, err
	}

	return newFacade(game), nil
}

func newFacade(game *game.Game) *gameFacade {
	checkService := service.NewCheckService(game)
	movementService := service.NewMovementService(game)
	checkMateService := service.NewCheckMateService(game)
	registerPlayerService := service.NewRegisterPlayerService(game)

	return &gameFacade{
		game:           game,
		registerPlayer: usecases.NewRegisterPlayer(registerPlayerService),
		move:           usecases.NewMove(game, movementService, checkService, checkMateService),
	}
}

func (f *gameFacade) RegisterPlayer(nick string) error {
	return f.registerPlayer.Execute(nick)
}

func (c *gameFacade) Render() {
	c.game.Render()
}

func (c *gameFacade) Board() board.Board {
	return c.game.Board
}

func (f *gameFacade) Move(iLine, iColumn, tLine, tColumn int) (err error) {
	return f.move.Execute(iLine, iColumn, tLine, tColumn)
}

func (f *gameFacade) Square(line, column int) *entity.Square {
	return f.game.Square(line, column)
}

func (f *gameFacade) Status() game.Status {
	return f.game.Status
}

func (f *gameFacade) GetState() GameState {
	return GameState{Board: f.game.Board.State()}
}
