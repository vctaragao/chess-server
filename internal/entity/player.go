package entity

import "github.com/google/uuid"

type Player struct {
	ID uuid.UUID
}

func NewPlayer() *Player {
	return &Player{
		ID: uuid.New(),
	}
}
