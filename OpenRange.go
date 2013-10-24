package tddbc

type OpenRange struct {
	lower int
	upper int
}

func NewOpenRange(lower, upper int) *OpenRange {
	return &OpenRange{lower, upper}
}
