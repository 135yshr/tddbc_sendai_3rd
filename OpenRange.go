package tddbc

import (
	"fmt"
)

type OpenRange struct {
	lower int
	upper int
}

func NewOpenRange(lower, upper int) *OpenRange {
	return &OpenRange{lower, upper}
}

func (r *OpenRange) String() string {
	return fmt.Sprintf("[%d,%d]", r.lower, r.upper)
}

func (r *OpenRange) Contains(target int) bool {
	return true
}
