package main

import (
	"encoding/json"
	"log"

	"github.com/vctaragao/chess-server/internal"
)

type (
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	MovementRequest struct {
		Action          string   `json:"action"`
		Result          string   `json:"result"`
		TargetPosition  Position `json:"target_pos"`
		InitialPosition Position `json:"initial_pos"`
	}
)

func main() {
	game := internal.NewGame()

	if err := game.RegisterPlayer(); err != nil {
		log.Fatal("unable to register first player", err)
	}

	if err := game.RegisterPlayer(); err != nil {
		log.Fatal("unable to register second player", err)
	}

	game.Render()

	var movRequest MovementRequest
	if err := json.Unmarshal(wPlayerMove(), &movRequest); err != nil {
		log.Fatal("unable to unmarshal movement request", err)
	}

	action := game.ParseAction(movRequest.Action)
	result := game.ParseResult(movRequest.Result)

	iSquare := game.GetSquare(movRequest.InitialPosition.Y, movRequest.InitialPosition.X)
	tSquare := game.GetSquare(movRequest.TargetPosition.Y, movRequest.TargetPosition.X)

	game.HandleMovement(iSquare, tSquare, action, result)

	game.Render()
}

func wPlayerMove() []byte {
	return []byte(`{
        "initial_pos": {
            "x": 6,
            "y": 7
        },
        "target_pos"{
            "x": 5,
            "y": 5
        },
        "result": "check",
        "action": "capture"
    }`)
}
