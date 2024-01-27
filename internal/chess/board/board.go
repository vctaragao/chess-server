package board

import (
	"fmt"

	"github.com/vctaragao/chess-server/internal/chess/entity"
)

type Board [8][8]*entity.Square

func NewBoard() Board {
	board := initializeBoard()
	board.UpdateAttackingSquares()

	return board
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
