package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vctaragao/chess/internal"
	"github.com/vctaragao/chess/internal/entity"
)

type square struct {
	x, y int
}

type Chess struct {
	board          [][]entity.Piece
	selectedSquare *square
}

func (c *Chess) Init() tea.Cmd {
	chessAdapter := &internal.ChessAdapter{}
	board := chessAdapter.NewGame()

	c.board = board
	c.selectedSquare = nil

	return nil
}

func (c Chess) View() string {
	if len(c.board) == 0 {
		return "Board vazio..."
	}

	displayBoard := make([][]string, 8)

	for i := range displayBoard {
		displayBoard[i] = make([]string, 8)
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			displayBoard[y][x] = c.board[y][x].String()
		}
	}
}

func main() {
	f, err := tea.LogToFile("chess.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("Chess game started")

	game := &Chess{}

	chessAdapter := &internal.ChessAdapter{}
	board := chessAdapter.NewGame()

	c.board = board
	c.selectedSquare = nil
}
