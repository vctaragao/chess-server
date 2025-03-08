package entity

import (
	"errors"
	"fmt"
)

type PieceType int

var ErrEmptySquare = errors.New("empty square")

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p PieceType) String() string {
	return []string{"P", "K", "B", "R", "Q", "k"}[p]
}

func NewPieceType(line, column int) (PieceType, error) {
	var pieceType PieceType
	var err error

	if line == 6 || line == 1 {
		pieceType = Pawn
		return pieceType, nil
	}

	switch column {
	case 0, 7:
		pieceType = Rook
	case 1, 6:
		pieceType = Knight
	case 2, 5:
		pieceType = Bishop
	case 3:
		pieceType = Queen
	case 4:
		pieceType = King
	default:
		err = fmt.Errorf("no piece type for line: %d and column: %d", line, column)
	}

	return pieceType, err
}

func PieceTypeFromString(p byte) (PieceType, error) {
	var pieceType PieceType
	var err error

	switch p {
	case 'P':
		pieceType = Pawn
	case 'K':
		pieceType = Knight
	case 'B':
		pieceType = Bishop
	case 'R':
		pieceType = Rook
	case 'Q':
		pieceType = Queen
	case 'k':
		pieceType = King
	case '_':
		err = ErrEmptySquare
	default:
		err = fmt.Errorf("unable to parse piece type from string: %s", string(p))
	}

	return pieceType, err
}
