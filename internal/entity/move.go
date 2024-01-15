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

func NewMovement(initialSquare, targetSquare Square, action Action, result Result) Movement {
	return Movement{
		InitialSquare: initialSquare,
		TargetSquare:  targetSquare,
		Action:        action,
		Result:        result,
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
	return m.TargetSquare.y
}

func (m *Movement) TargetX() int {
	return m.TargetSquare.x
}

func (m *Movement) InitialY() int {
	return m.InitialSquare.y
}

func (m *Movement) InitialX() int {
	return m.InitialSquare.x
}

func (m *Movement) GetPiece() *Piece {
	return m.InitialSquare.piece
}
