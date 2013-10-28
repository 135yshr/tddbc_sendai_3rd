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
	return fmt.Sprintf("(%d,%d)", r.lower, r.upper)
}

func (r *OpenRange) Contains(target int) bool {
	return r.lower < target && target < r.upper
}

func (r *OpenRange) Equal(target OpenRange) bool {
	return r.lower == target.lower && r.upper == target.upper
}

func (r *OpenRange) IsConnectedTo(target OpenRange) bool {
	return true
}
