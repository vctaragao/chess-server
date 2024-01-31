package service

import (
	"slices"

	"github.com/vctaragao/chess-server/internal/chess/entity"
	"github.com/vctaragao/chess-server/internal/chess/game"
	"github.com/vctaragao/chess-server/internal/chess/helper"
)

type CheckMateService struct {
	*game.Game
	m       *entity.Movement
	kSquare *entity.Square
}

func NewCheckMateService(g *game.Game) *CheckMateService {
	return &CheckMateService{
		Game: g,
	}
}

func (s *CheckMateService) HandleCheckMate(m *entity.Movement, kSquare *entity.Square) {
	s.m = m
	s.kSquare = kSquare

	// check if king can move or if any piece can block the attack or capture the attacking piece
	if s.canKingMove() || s.canPieceBeBlocked() || s.canPieceBeCaptured() {
		return
	}

	s.Status = game.CheckMate
}

func (s *CheckMateService) canKingMove() bool {
	piece := s.m.GetPiece()

	kColor := s.kSquare.Piece.Color
	kingPossibleSquares := s.kSquare.Piece.AttackingSquares

	attackingSquares := s.GetAllAttackingSquares(piece.Color)

	for _, square := range kingPossibleSquares {
		// if square is not empty
		// and has a piece of the same color
		// or it has a piece of a different color that it's protected by another piece
		if !square.IsEmpty() && (square.Piece.HasColor(kColor) || square.Piece.IsProteced()) {
			continue
		}

		if !slices.Contains(attackingSquares, square) {
			return true
		}
	}

	return false
}

func (s *CheckMateService) canPieceBeBlocked() bool {
	piece := s.m.GetPiece()

	if piece.Is(entity.Knight) || piece.Is(entity.Pawn) {
		return false
	}

	// get all checkingSquares
	var checkingSquares []*entity.Square

	switch piece.PieceType {
	case entity.Bishop:
		checkingSquares = s.getBishopCheckingSquares()
	case entity.Rook:
		checkingSquares = s.getRookCheckingSquares()
	case entity.Queen:
		checkingSquares = s.getQueenCheckingSquares()
	}

	// get all squares that enemy can attack
	color := helper.White
	if piece.IsWhite() {
		color = helper.Black
	}

	attackinSquares := s.GetAllAttackingSquares(color)

	mappedSquares := make(map[helper.Position]struct{})
	for _, square := range attackinSquares {
		mappedSquares[square.Position] = struct{}{}
	}

	// check if any of the checking squares can be blocked by an eney piece
	for _, square := range checkingSquares {
		if _, ok := mappedSquares[square.Position]; ok {
			return true
		}
	}

	return false
}

func (s *CheckMateService) getBishopCheckingSquares() []*entity.Square {
	bSquare := s.m.TargetSquare

	// get all squares between bishop and king
	var squares []*entity.Square
	// down right
	if s.kSquare.Y > bSquare.Y && s.kSquare.X > bSquare.X {
		for y, x := bSquare.Y+1, bSquare.X+1; y < s.kSquare.Y && x < s.kSquare.X; y, x = y+1, x+1 {
			squares = append(squares, s.Board[y][x])
		}

		return squares
	}

	// down left
	if s.kSquare.Y > bSquare.Y && s.kSquare.X < bSquare.X {
		for y, x := bSquare.Y+1, bSquare.X-1; y < s.kSquare.Y && x > s.kSquare.X; y, x = y+1, x-1 {
			squares = append(squares, s.Board[y][x])
		}

		return squares
	}

	// up right
	if s.kSquare.Y < bSquare.Y && s.kSquare.X > bSquare.X {
		for y, x := bSquare.Y-1, bSquare.X+1; y > s.kSquare.Y && x < s.kSquare.X; y, x = y-1, x+1 {
			squares = append(squares, s.Board[y][x])
		}

		return squares
	}

	// up left
	if s.kSquare.Y < bSquare.Y && s.kSquare.X < bSquare.X {
		for y, x := bSquare.Y-1, bSquare.X-1; y > s.kSquare.Y && x > s.kSquare.X; y, x = y-1, x-1 {
			squares = append(squares, s.Board[y][x])
		}

		return squares
	}

	return nil
}

func (s *CheckMateService) getRookCheckingSquares() []*entity.Square {
	rSquare := s.m.TargetSquare

	// get all squares between rook and king
	var squares []*entity.Square
	// down
	if s.kSquare.Y > rSquare.Y {
		for y := rSquare.Y + 1; y < s.kSquare.Y; y++ {
			squares = append(squares, s.Board[y][rSquare.X])
		}

		return squares
	}

	// up
	if s.kSquare.Y < rSquare.Y {
		for y := rSquare.Y - 1; y > s.kSquare.Y; y-- {
			squares = append(squares, s.Board[y][rSquare.X])
		}

		return squares
	}

	// right
	if s.kSquare.X > rSquare.X {
		for x := rSquare.X + 1; x < s.kSquare.X; x++ {
			squares = append(squares, s.Board[rSquare.Y][x])
		}

		return squares
	}

	// left
	if s.kSquare.X < rSquare.X {
		for x := rSquare.X - 1; x > s.kSquare.X; x-- {
			squares = append(squares, s.Board[rSquare.Y][x])
		}

		return squares
	}

	return nil
}

func (s *CheckMateService) getQueenCheckingSquares() []*entity.Square {
	qSquare := s.m.TargetSquare

	if qSquare.Y == s.kSquare.Y || qSquare.X == s.kSquare.X {
		return s.getRookCheckingSquares()
	}

	return s.getBishopCheckingSquares()
}

func (s *CheckMateService) canPieceBeCaptured() bool {
	// get all squares that enemy can attack
	piece := s.m.GetPiece()

	color := helper.White
	if piece.IsWhite() {
		color = helper.Black
	}

	attackinSquares := s.GetAllAttackingSquares(color)

	// check if any of the checking squares can be captured by an eney piece
	return slices.Contains(attackinSquares, piece.Square)
}
