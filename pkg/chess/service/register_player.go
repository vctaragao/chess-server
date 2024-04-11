package service

import (
	"errors"

	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/game"
)

var ErrMaxPlayer = errors.New("Game already full of players")

type RegisterPlayerService struct {
	*game.Game
}

func NewRegisterPlayerService(g *game.Game) *RegisterPlayerService {
	return &RegisterPlayerService{
		Game: g,
	}
}

func (s *RegisterPlayerService) Execute(nick string) error {
	if s.WPlayer != nil && s.BPlayer != nil {
		return ErrMaxPlayer
	}

	if s.WPlayer != nil {
		s.BPlayer = entity.NewPlayer(nick)
	} else {
		s.WPlayer = entity.NewPlayer(nick)
	}

	return nil
}
