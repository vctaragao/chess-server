package usecases

import "github.com/vctaragao/chess/pkg/chess/service"

type RegisterPlayer struct {
	RegisterPlayerService *service.RegisterPlayerService
}

func NewRegisterPlayer(r *service.RegisterPlayerService) *RegisterPlayer {
	return &RegisterPlayer{
		RegisterPlayerService: r,
	}
}

func (r *RegisterPlayer) Execute(nick string) error {
	return r.RegisterPlayerService.Execute(nick)
}
