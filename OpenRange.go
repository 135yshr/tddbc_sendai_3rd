package tddbc

type OpenRange struct {
	lower int
}

func NewOpenRange(lower, upper int) *OpenRange {
	return &OpenRange{lower}
}
