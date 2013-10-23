package tddbc

type ClosedRange struct {
	lower int
	upper int
}

func NewClosedRange(lower, upper int) *ClosedRange {
	return &ClosedRange{lower, upper}
}
