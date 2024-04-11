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

func NewBoard() Board {
	board := initializeBoard()
	board.UpdateAttackingSquares()

	return board
}

func NewBoardFromString(board string) Board {
	b := Board{}

	board = strings.ReplaceAll(board, "\n", "")

	reader := strings.NewReader(board)
	y, x := 0, 0

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

		color := helper.White
		if piece[0] == 'B' {
			color = helper.Black
		}

		pType := entity.None
		switch piece[1] {
		case 'R':
			pType = entity.Rook
		case 'B':
			pType = entity.Bishop
		case 'K':
			pType = entity.Knight
		case 'Q':
			pType = entity.Queen
		case 'k':
			pType = entity.King
		case 'P':
			pType = entity.Pawn
		}

		p := entity.NewEmptyPiece()
		if pType != entity.None {
			p = entity.NewPiece(color, pType)
		}

		square := entity.NewSquare(y, x)
		square.SetPiece(p)

		b[y][x] = square
		x++

		if x == 8 {
			x = 0
			y++
		}
	}

	return b
}

func initializeBoard() Board {
	initialBoard := Board{}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			p := entity.NewEmptyPiece()
			square := entity.NewSquare(y, x)

			if !isPieceLine(y) {
				initialBoard.insert(y, x, square, p)
				continue
			}

			color := colorByLine(y)
			p = entity.NewPiece(color, entity.Pawn)

			if y == 0 || y == 7 {
				var pType entity.PieceType
				switch x {
				case 0, 7:
					pType = entity.Rook
				case 1, 6:
					pType = entity.Knight
				case 2, 5:
					pType = entity.Bishop
				case 3:
					pType = entity.Queen
				case 4:
					pType = entity.King
				}

				p.PieceType = pType
			}

			initialBoard.insert(y, x, square, p)
		}
	}

	return initialBoard
}

func (b *Board) insert(y, x int, square *entity.Square, piece *entity.Piece) {
	square.SetPiece(piece)

	if !piece.IsNull {
		piece.Square = square
	}

	b[y][x] = square
}

func (b Board) GetPiece(y, x int) *entity.Piece {
	return b[y][x].Piece
}

func (b Board) Show() string {
	strBoard := "\n"
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			s := b[y][x]
			if s.IsEmpty() {
				strBoard += fmt.Sprintf(" %s ", entity.None)
				continue
			}

			strBoard += fmt.Sprintf(" %s%s ", s.Piece.Color, s.Piece.PieceType)
		}
		strBoard += "\n"
	}

	return strBoard
}

func (b Board) State() [][]string {
	log.Println("State")
	state := make([][]string, 8)
	for i := range state {
		state[i] = make([]string, 8)
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			log.Println("Square", y, x)
			s := b[y][x]
			if s.IsEmpty() {
				log.Println("Empty")
				state[y][x] = string(entity.None)
				continue
			}

			log.Println("Piece", s.Piece)
			strPiece := fmt.Sprintf("%s%s", s.Piece.Color, s.Piece.PieceType)
			log.Println("strPiece", strPiece)

			state[y][x] = fmt.Sprintf("%s%s", s.Piece.Color, s.Piece.PieceType)
			log.Println("Piece inserted into State")
		}
	}

	log.Println("Finished state", state)

	return state
}
