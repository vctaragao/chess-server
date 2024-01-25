package helper

type Color int

const (
	White Color = 0
	Black Color = 1
)

func (c Color) String() string {
	return []string{"W", "B"}[c]
}
