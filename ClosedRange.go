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
	return r.lower == target.lower && r.upper == target.upper
}

func (r *ClosedRange) IsConnectedTo(target ClosedRange) bool {
	return target.lower < r.upper && r.lower < target.upper
}
