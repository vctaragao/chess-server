package entity

import "github.com/google/uuid"

type Player struct {
	ID     uuid.UUID
	Nick   string
	Points int
}

func NewPlayer(nick string) *Player {
	return &Player{
		ID:     uuid.New(),
		Nick:   nick,
		Points: 0,
	}
}
