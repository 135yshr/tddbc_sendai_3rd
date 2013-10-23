package tddbc

type ClosedRange struct {
	lower int
	upper int
}

func NewClosedRange(lower, upper int) *ClosedRange {
	if upper < lower {
		return nil
	}
	return &ClosedRange{lower, upper}
}
