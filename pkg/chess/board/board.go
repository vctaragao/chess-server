package board

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/helper"
)

type Board [8][8]*entity.Square

func NewBoard() (Board, error) {
	board, err := initializeBoard()
	if err != nil {
		return board, fmt.Errorf("initializing board: %w", err)
	}
	board.UpdateAttackingSquares()

	return board, nil
}

func NewBoardFromString(boardStr string) (Board, error) {
	board, err := initializeBoardFromString(boardStr)
	if err != nil {
		return board, fmt.Errorf("initializing board from string: %w", err)
	}

	board.UpdateAttackingSquares()

	return board, nil
}

func initializeBoardFromString(boardStr string) (Board, error) {
	board := Board{}

	boardStr = strings.TrimSpace(boardStr)
	boardStr = strings.ReplaceAll(boardStr, "\n", "  ")

	reader := strings.NewReader(boardStr)
	line, column := 0, 0

	for {
		piece := make([]byte, 2)
		_, err := reader.Read(piece)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		if piece[0] == ' ' {
			continue
		}

		square := entity.NewSquare(line, column)

		color := helper.White
		if piece[0] == 'B' {
			color = helper.Black
		}

		pType, err := entity.PieceTypeFromString(piece[1])
		if err != nil && err != entity.ErrEmptySquare {
			return board, err
		}

		var p *entity.Piece
		if err != entity.ErrEmptySquare {
			p = entity.NewPiece(color, pType, square)
		}

		square.SetPiece(p)

		board[line][column] = square
		column++

		if column == 8 {
			column = 0
			line++
		}
	}

	return board, nil

}

func initializeBoard() (Board, error) {
	initialBoard := Board{}
	for line := 0; line < 8; line++ {
		for column := 0; column < 8; column++ {
			var piece *entity.Piece
			square := entity.NewSquare(line, column)

			if isPieceLine(line) {
				color := colorByLine(line)
				pieceType, err := entity.NewPieceType(line, column)
				if err != nil {
					return initialBoard, err
				}

				piece = entity.NewPiece(color, pieceType, square)
			}

			square.SetPiece(piece)

			initialBoard.insert(line, column, square)
		}
	}

	return initialBoard, nil
}

func (b *Board) insert(y, x int, square *entity.Square) {
	b[y][x] = square
}

func (b Board) GetPiece(y, x int) *entity.Piece {
	return b[y][x].Piece
}

func (b Board) Show() string {
	strBoard := "\n"
	for line := 0; line < 8; line++ {
		for column := 0; column < 8; column++ {
			s := b[line][column]
			if s.IsEmpty() {
				strBoard += " __ "
				continue
			}

			strBoard += fmt.Sprintf(" %s%s ", s.Piece.Color, s.Piece.PieceType)
		}
		strBoard += "\n"
	}

	return strBoard
}

func (b Board) State() [][]string {
	state := make([][]string, 8)
	for i := range state {
		state[i] = make([]string, 8)
	}

	for line := 0; line < 8; line++ {
		for column := 0; column < 8; column++ {
			s := b[line][column]
			if s.IsEmpty() {
				state[line][column] = "   "
				continue
			}

			state[line][column] = fmt.Sprintf("%s%s", s.Piece.Color, s.Piece.PieceType)
		}
	}

	return state
}
