package game

type Status int

const (
	None Status = iota
	Check
	CheckMate
)

func (s Status) String() string {
	return []string{"None", "Check", "CheckMate"}[s]
}
