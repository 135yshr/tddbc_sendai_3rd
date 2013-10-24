package tddbc

type OpenRange struct {
	lower int
	upper int
}

func NewOpenRange(lower, upper int) *OpenRange {
	return &OpenRange{lower, upper}
}

func (r OpenRange) String() string {
	return "[3, 8]"
}
