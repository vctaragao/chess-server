package board

import (
	"slices"

	"github.com/vctaragao/chess/pkg/chess/entity"
	"github.com/vctaragao/chess/pkg/chess/helper"
)

func (b *Board) UpdateAttackingSquares() {
	b.ClearProtecedBy()

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			square := b[y][x]
			if square.IsEmpty() {
				continue
			}

			piece := square.Piece
			piece.AttackingSquares = b.getAttackingSquares(piece)
		}
	}
}

func (b *Board) ClearProtecedBy() {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			square := b[y][x]
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
	y, x := p.Square.Y, p.Square.X

	// up
	if y-1 >= 0 {
		targetSquare := b[y-1][x]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// down
	if y+1 <= 7 {
		targetSquare := b[y+1][x]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// left
	if x-1 >= 0 {
		targetSquare := b[y][x-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// right
	if x+1 <= 7 {
		targetSquare := b[y][x+1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// up-left
	if y-1 >= 0 && x-1 >= 0 {
		targetSquare := b[y-1][x-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// up-right
	if y-1 >= 0 && x+1 <= 7 {
		targetSquare := b[y-1][x+1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// down-left
	if y+1 <= 7 && x-1 >= 0 {
		targetSquare := b[y+1][x-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, p, p.Color)
	}

	// down-right
	if y+1 <= 7 && x+1 <= 7 {
		targetSquare := b[y+1][x+1]
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
	y, x := piece.Square.Y, piece.Square.X

	if y-2 >= 0 && x-1 >= 0 {
		tSquare := b[y-2][x-1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y-2 >= 0 && x+1 <= 7 {
		tSquare := b[y-2][x+1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y-1 >= 0 && x-2 >= 0 {
		tSquare := b[y-1][x-2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y-1 >= 0 && x+2 <= 7 {
		tSquare := b[y-1][x+2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y+1 <= 7 && x-2 >= 0 {
		tSquare := b[y+1][x-2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y+1 <= 7 && x+2 <= 7 {
		tSquare := b[y+1][x+2]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y+2 <= 7 && x-1 >= 0 {
		tSquare := b[y+2][x-1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	if y+2 <= 7 && x+1 <= 7 {
		tSquare := b[y+2][x+1]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)
	}

	return squares
}

func (b *Board) getBishopAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.Square.Y, piece.Square.X

	// diagonal up-left
	for up, left := y-1, x-1; up >= 0 && left >= 0; up, left = up-1, left-1 {
		tSquare := b[up][left]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[up][left].IsEmpty() {
			break
		}
	}

	// diagonal up-right
	for up, right := y-1, x+1; up >= 0 && right <= 7; up, right = up-1, right+1 {
		tSquare := b[up][right]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[up][right].IsEmpty() {
			break
		}
	}

	// diagonal down-left
	for down, left := y+1, x-1; down <= 7 && left >= 0; down, left = down+1, left-1 {
		tSquare := b[down][left]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[down][left].IsEmpty() {
			break
		}
	}

	// diagonal down-right
	for down, right := y+1, x+1; down <= 7 && right <= 7; down, right = down+1, right+1 {
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
	y, x := piece.Square.Y, piece.Square.X

	// up
	for down := y - 1; down >= 0; down-- {
		tSquare := b[down][x]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[down][x].IsEmpty() {
			break
		}
	}

	// down
	for up := y + 1; up <= 7; up++ {
		tSquare := b[up][x]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[up][x].IsEmpty() {
			break
		}
	}

	// left
	for left := x - 1; left >= 0; left-- {
		tSquare := b[y][left]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[y][left].IsEmpty() {
			break
		}
	}

	// right
	for right := x + 1; right <= 7; right++ {
		tSquare := b[y][right]
		squares = append(squares, tSquare)
		b.updateProtectedBy(tSquare, piece, piece.Color)

		if !b[y][right].IsEmpty() {
			break
		}
	}

	return squares
}

func (b *Board) getPawnAttackingSquares(piece *entity.Piece) []*entity.Square {
	var squares []*entity.Square
	y, x := piece.Square.Y, piece.Square.X

	if piece.IsWhite() {
		if y-1 >= 0 && x-1 >= 0 {
			targetSquare := b[y-1][x-1]
			squares = append(squares, targetSquare)
			b.updateProtectedBy(targetSquare, piece, helper.White)
		}

		if y-1 >= 0 && x+1 <= 7 {
			targetSquare := b[y-1][x+1]
			squares = append(squares, targetSquare)
			b.updateProtectedBy(targetSquare, piece, helper.White)
		}

		return squares
	}

	if y+1 <= 7 && x-1 >= 0 {
		targetSquare := b[y+1][x-1]
		squares = append(squares, targetSquare)
		b.updateProtectedBy(targetSquare, piece, helper.Black)
	}

	if y+1 <= 7 && x+1 <= 7 {
		targetSquare := b[y+1][x+1]
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
