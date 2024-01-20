package entity

type (
	Action int
	Result int

	Movement struct {
		InitialSquare Square
		TargetSquare  Square
		Action        Action
		Result        Result
	}
)

const (
	Move Action = iota
	Capture

	Empty Result = iota
	Check
	CheckMate
)

func NewMovement(initialSquare, targetSquare Square) Movement {
	return Movement{
		InitialSquare: initialSquare,
		TargetSquare:  targetSquare,
	}
}

func NewActionFromString(action string) Action {
	if action == "capture" {
		return Capture
	}

	return Move
}

func NewResultFromString(result string) Result {
	switch result {
	case "check":
		return Check
	case "check_mate":
		return CheckMate
	default:
		return Empty
	}
}

func (m *Movement) TargetY() int {
	return m.TargetSquare.Y
}

func (m *Movement) TargetX() int {
	return m.TargetSquare.X
}

func (m *Movement) InitialY() int {
	return m.InitialSquare.Y
}

func (m *Movement) InitialX() int {
	return m.InitialSquare.X
}

func (m *Movement) GetPiece() *Piece {
	return m.InitialSquare.piece
}

func (m *Movement) GetTargetPiece() *Piece {
	return m.TargetSquare.piece
}

func (m *Movement) IsValid() bool {
	if !m.TargetSquare.IsEmpty() && m.TargetSquare.piece.color == m.InitialSquare.piece.color {
		return false
	}

	return true
}

func (m *Movement) IsCapture() bool {
	return !m.TargetSquare.IsEmpty() && m.TargetSquare.piece.color != m.InitialSquare.piece.color
}
