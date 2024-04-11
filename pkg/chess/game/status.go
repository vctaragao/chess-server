package game

type Status int

const (
	None Status = iota
	Check
	CheckMate
)

func (s *Status) String() string {
	switch *s {
	case None:
		return "None"
	case Check:
		return "Check"
	case CheckMate:
		return "CheckMate"
	}

	return ""
}
