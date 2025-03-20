package board

import (
	"slices"

	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/helper"
)

func (b *Board) UpdateAttackingSquares() {
	b.ClearProtecedBy()

	for line := 0; line < 8; line++ {
		for column := 0; column < 8; column++ {
			square := b[line][column]
			if square.IsEmpty() {
				continue
			}

			piece := square.Piece
			piece.AttackingSquares = b.getAttackingSquares(piece)
		}
	}
}

func (b *Board) ClearProtecedBy() {
	for line := 0; line < 8; line++ {
		for column := 0; column < 8; column++ {
			square := b[line][column]
			if square.IsEmpty() {
				continue
			}

			square.Piece.ProtecedBy = nil
		}
	}
}

func (b *Board) getAttackingSquares(p *entity.Piece) []*entity.Square {
	switch p.PieceType {
	case entity.Pawn:
		return b.getPawnAttackingSquares(p)
	case entity.Knight:
		return b.getKnightAttackingSquares(p)
	case entity.Bishop:
		return b.getBishopAttackingSquares(p)
	case entity.Rook:
		return b.getRookAttackingSquares(p)
	case entity.Queen:
		return b.getQueenAttackingSquares(p)
	case entity.King:
		return b.getKingAttackingSquares(p)
	default:
		return []*entity.Square{}
	}
}

func (b *Board) getKingAttackingSquares(p *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	line, column := p.Square.Line, p.Square.Column

	// up
	if line-1 >= 0 {
		targetSquare := b[line-1][column]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// down
	if line+1 <= 7 {
		targetSquare := b[line+1][column]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// left
	if column-1 >= 0 {
		targetSquare := b[line][column-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// right
	if column+1 <= 7 {
		targetSquare := b[line][column+1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// up-left
	if line-1 >= 0 && column-1 >= 0 {
		targetSquare := b[line-1][column-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// up-right
	if line-1 >= 0 && column+1 <= 7 {
		targetSquare := b[line-1][column+1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// down-left
	if line+1 <= 7 && column-1 >= 0 {
		targetSquare := b[line+1][column-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// down-right
	if line+1 <= 7 && column+1 <= 7 {
		targetSquare := b[line+1][column+1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	return squares
}

func (b *Board) getQueenAttackingSquares(p *entity.Piece) []*entity.Square {
	return append(b.getRookAttackingSquares(p), b.getBishopAttackingSquares(p)...)
}

func (b *Board) getKnightAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	line, column := piece.Square.Line, piece.Square.Column

	if line-2 >= 0 && column-1 >= 0 {
		tSquare := b[line-2][column-1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line-2 >= 0 && column+1 <= 7 {
		tSquare := b[line-2][column+1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line-1 >= 0 && column-2 >= 0 {
		tSquare := b[line-1][column-2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line-1 >= 0 && column+2 <= 7 {
		tSquare := b[line-1][column+2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line+1 <= 7 && column-2 >= 0 {
		tSquare := b[line+1][column-2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line+1 <= 7 && column+2 <= 7 {
		tSquare := b[line+1][column+2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line+2 <= 7 && column-1 >= 0 {
		tSquare := b[line+2][column-1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if line+2 <= 7 && column+1 <= 7 {
		tSquare := b[line+2][column+1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	return squares
}

func (b *Board) getBishopAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	line, column := piece.Square.Line, piece.Square.Column

	// diagonal up-left
	for up, left := line-1, column-1; up >= 0 && left >= 0; up, left = up-1, left-1 {
		tSquare := b[up][left]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[up][left].IsEmpty() {
			break
		}
	}

	// diagonal up-right
	for up, right := line-1, column+1; up >= 0 && right <= 7; up, right = up-1, right+1 {
		tSquare := b[up][right]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[up][right].IsEmpty() {
			break
		}
	}

	// diagonal down-left
	for down, left := line+1, column-1; down <= 7 && left >= 0; down, left = down+1, left-1 {
		tSquare := b[down][left]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[down][left].IsEmpty() {
			break
		}
	}

	// diagonal down-right
	for down, right := line+1, column+1; down <= 7 && right <= 7; down, right = down+1, right+1 {
		tSquare := b[down][right]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[down][right].IsEmpty() {
			break
		}
	}

	return squares
}

func (b *Board) getRookAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	line, column := piece.Square.Line, piece.Square.Column

	// up
	for down := line - 1; down >= 0; down-- {
		tSquare := b[down][column]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[down][column].IsEmpty() {
			break
		}
	}

	// down
	for up := line + 1; up <= 7; up++ {
		tSquare := b[up][column]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[up][column].IsEmpty() {
			break
		}
	}

	// left
	for left := column - 1; left >= 0; left-- {
		tSquare := b[line][left]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[line][left].IsEmpty() {
			break
		}
	}

	// right
	for right := column + 1; right <= 7; right++ {
		tSquare := b[line][right]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[line][right].IsEmpty() {
			break
		}
	}

	return squares
}

func (b *Board) getPawnAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	line, column := piece.Square.Line, piece.Square.Column

	if piece.IsWhite() {
		if line-1 >= 0 && column-1 >= 0 {
			targetSquare := b[line-1][column-1]
			squares = append(squares, targetSquare)
			b.updateProtectedBy(targetSquare, piece, helper.White)
		}

		if line-1 >= 0 && column+1 <= 7 {
			targetSquare := b[line-1][column+1]
			squares = append(squares, targetSquare)
			b.updateProtectedBy(targetSquare, piece, helper.White)
		}

		return squares
	}

	if line+1 <= 7 && column-1 >= 0 {
		targetSquare := b[line+1][column-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, piece, helper.Black)
	}

	if line+1 <= 7 && column+1 <= 7 {
		targetSquare := b[line+1][column+1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, piece, helper.Black)
	}

	return squares
}

func (b *Board) updateProtectedBy(targetSquare *entity.Square, piece *entity.Piece, color helper.Color) {
	if !targetSquare.IsEmpty() && targetSquare.Piece.HasColor(color) {
		targetSquare.Piece.ProtecedBy = piece
	}
}

func isPieceLine(y int) bool {
	return slices.Contains([]int{0, 1, 6, 7}, y)
}

func colorByLine(y int) helper.Color {
	color := helper.White
	if y == 1 || y == 0 {
		color = helper.Black
	}

	return color
}
