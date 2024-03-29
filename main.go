package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/vctaragao/chess-server/internal/chess"
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
	game := chess.NewGame()

	if err := game.RegisterPlayer("Victor"); err != nil {
		log.Fatal("unable to register first player", err)
	}

	if err := game.RegisterPlayer("Pedro"); err != nil {
		log.Fatal("unable to register second player", err)
	}

	game.Render()

	moves := []func() []byte{
		wFirstPawnMove,
		bFirstMove,
		wSecondPawnMove,
		bSecondMove,
	}

	for _, move := range moves {
		var m MovementRequest
		if err := json.Unmarshal(move(), &m); err != nil {
			log.Fatal("unable to unmarshal movement request", err)
		}

		iSquare := game.GetSquare(m.InitialPosition.Y, m.InitialPosition.X)
		tSquare := game.GetSquare(m.TargetPosition.Y, m.TargetPosition.X)
		if err := game.Move(iSquare, tSquare); err != nil {
			log.Println(err, m)
		}

		game.Render()
	}

	fmt.Println("Status: ", game.GetStatus())
}

func wFirstPawnMove() []byte {
	return []byte(`{
        "initial_pos": {
            "x": 5,
            "y": 6
        },
        "target_pos": {
            "x": 5,
            "y": 5
        }
    }`)
}

func wSecondPawnMove() []byte {
	return []byte(`{
        "initial_pos": {
            "x": 6,
            "y": 6
        },
        "target_pos": {
            "x": 6,
            "y": 4
        }
    }`)
}

func bFirstMove() []byte {
	return []byte(`{
        "initial_pos": {
            "x": 4,
            "y": 1
        },
        "target_pos": {
            "x": 4,
            "y": 2
        }
    }`)
}

func bSecondMove() []byte {
	return []byte(`{
        "initial_pos": {
            "x": 3,
            "y": 0
        },
        "target_pos": {
            "x": 7,
            "y": 4
        }
    }`)
}
