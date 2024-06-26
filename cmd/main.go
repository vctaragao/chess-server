package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vctaragao/chess/internal"
	"github.com/vctaragao/chess/internal/entity"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

func (c *Chess) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s := msg.String(); s == "ctrl+c" || s == "q" || s == "esc" {
			return c, tea.Quit
		}
	}

	return c, nil
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

	re := lipgloss.NewRenderer(os.Stdout)
	labelStyle := re.NewStyle().Foreground(lipgloss.Color("241"))

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		Rows(displayBoard...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return lipgloss.NewStyle().Padding(0, 1)
		})

	ranks := labelStyle.Render(strings.Join([]string{" A", "B", "C", "D", "E", "F", "G", "H  "}, "   "))
	files := labelStyle.Render(strings.Join([]string{" 1", "2", "3", "4", "5", "6", "7", "8 "}, "\n\n "))

	return lipgloss.JoinVertical(lipgloss.Right, lipgloss.JoinHorizontal(lipgloss.Center, files, t.Render()), ranks) + "\n"
}

func main() {
	f, err := tea.LogToFile("chess.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Println("Chess game started")

	game := &Chess{}
	p := tea.NewProgram(game, tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Chess game finished")
}
