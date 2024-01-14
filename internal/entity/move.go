package entity

type (
	Action int
	Result int

	Movement struct {
		initialSquare Square
		targetSquare  Square
		action        Action
		result        Result
	}
)

const (
	Default Action = iota
	Capture

	Empty Result = iota
	Check
	CheckMate
)

func NewMovement(initialSquare, targetSquare Square, action Action, result Result) Movement {
	return Movement{
		initialSquare: initialSquare,
		targetSquare:  targetSquare,
		action:        action,
		result:        result,
	}
}

func NewActionFromString(action string) Action {
	if action == "capture" {
		return Capture
	}

	return Default
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
	return m.targetSquare.y
}

func (m *Movement) TargetX() int {
	return m.targetSquare.x
}

func (m *Movement) GetPiece() *Piece {
	return m.initialSquare.piece
}
