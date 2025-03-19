package helper

import "errors"

var ErrInvalidColor = errors.New("invalid color")

type Color int

const (
	White Color = 0
	Black Color = 1
)

func (c Color) String() string {
	return []string{"W", "B"}[c]
}

func ColorFromStr(color byte) (Color, error) {
	if color == 'W' {
		return White, nil
	}

	if color == 'B' {
		return Black, nil
	}

	return White, ErrInvalidColor
}
