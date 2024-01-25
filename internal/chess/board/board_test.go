package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	expectedBoard := `
 BR  BK  BB  BQ  Bk  BB  BK  BR 
 BP  BP  BP  BP  BP  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  WP  WP  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`

	board := NewBoard()
	assert.Equal(t, expectedBoard, board.Show())
}
