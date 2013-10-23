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

func (r *ClosedRange) Contains(target int) bool {
	return r.lower <= target && target <= r.upper
}

func (r *ClosedRange) Equal(target ClosedRange) bool {
	return true
}
