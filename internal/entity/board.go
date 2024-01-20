package entity

import (
	"fmt"
	"slices"
)

type Board [8][8]*Square

func NewBoard() Board {
	return initializeBoard()
}

func initializeBoard() Board {
	initialBoard := Board{}
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := NewEmptyPiece()
			square := NewSquare(y, x)

			if !isPieceLine(y) {
				initialBoard.insert(y, x, square, piece)
				continue
			}

			color := colorByLine(y)
			piece = NewPiece(color, Pawn)

			if y == 0 || y == 7 {
				var pType PieceType
				switch x {
				case 0, 7:
					pType = Rook
				case 1, 6:
					pType = Knight
				case 2, 5:
					pType = Bishop
				case 3:
					pType = Queen
				case 4:
					pType = King
				}

				piece.pieceType = pType
			}

			initialBoard.insert(y, x, square, piece)
		}
	}

	return initialBoard
}

func isPieceLine(y int) bool {
	return slices.Contains([]int{0, 1, 6, 7}, y)
}

func colorByLine(y int) Color {
	color := White
	if y == 1 || y == 0 {
		color = Black
	}

	return color
}

func (b *Board) insert(y, x int, square *Square, piece *Piece) {
	if !piece.isNull {
		square.SetPiece(piece)
		piece.SetSquare(square)
	}

	b[y][x] = square
}

func (b Board) GetPiece(y, x int) *Piece {
	return b[y][x].piece
}

func (b Board) Show() string {
	strBoard := "\n"
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			s := b[y][x]
			if s.empty {
				strBoard += fmt.Sprintf(" %s ", None)
				continue
			}

			strBoard += fmt.Sprintf(" %s%s ", s.piece.color, s.piece.GetType())
		}
		strBoard += "\n"
	}

	return strBoard
}
