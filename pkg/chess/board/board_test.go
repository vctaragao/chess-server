package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vctaragao/chess/pkg/chess/entity"
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
	board, err := NewBoard()
	assert.NoError(t, err)

	assert.Equal(t, expectedBoard, board.Show())
}

func TestNewBoardFromString(t *testing.T) {
	inputBoard := `
BR  BK  BB  BQ  Bk  BB  BK  BR
BP  BP  BP  BP  BP  BP  BP  BP
__  __  __  __  __  __  __  __
__  __  __  __  __  __  __  __
__  __  __  __  __  __  __  __
__  __  __  __  __  __  __  __
WP  WP  WP  WP  WP  WP  WP  WP
WR  WK  WB  WQ  Wk  WB  WK  WR
 `

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

	board, err := NewBoardFromString(inputBoard)
	require.NoError(t, err)

	assert.Equal(t, expectedBoard, board.Show())
}

func TestNewBoardFromStringCustom(t *testing.T) {
	initialBoard := `
 BR  BK  BB  __  Bk  BB  BK  BR 
 BP  BP  BP  BP  BP  BP  BP  BP 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  BQ  __  __ 
 __  __  __  __  __  __  __  __ 
 __  __  __  __  __  __  __  __ 
 WP  WP  WP  WP  __  WP  WP  WP 
 WR  WK  WB  WQ  Wk  WB  WK  WR 
`

	board, err := NewBoardFromString(initialBoard)
	require.NoError(t, err)

	assert.Equal(t, initialBoard, board.Show())
}

func TestProtectedBy(t *testing.T) {
	board, err := NewBoard()
	assert.NoError(t, err)

	shouldBeProtected := []*entity.Square{
		// BPieces
		board[0][1],
		board[0][2],
		board[0][3],
		board[0][4],
		board[0][5],
		board[0][6],

		// BPawns
		board[1][0],
		board[1][1],
		board[1][2],
		board[1][3],
		board[1][4],
		board[1][5],
		board[1][6],

		// WPieces
		board[7][1],
		board[7][2],
		board[7][3],
		board[7][4],
		board[7][5],
		board[7][6],

		// WPawns
		board[6][0],
		board[6][1],
		board[6][2],
		board[6][3],
		board[6][4],
		board[6][5],
		board[6][6],
	}

	shouldNotBeProteced := []*entity.Square{
		board[0][0],
		board[0][7],

		board[7][0],
		board[7][7],
	}

	for _, square := range shouldBeProtected {
		assert.True(t, square.Piece.IsProteced())
	}

	for _, square := range shouldNotBeProteced {
		assert.False(t, square.Piece.IsProteced())
	}
}
