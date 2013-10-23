package tddbc

import (
	"testing"
)

func Test正常_インスタンスが作成できる(t *testing.T) {
	target := NewClosedRange(3, 8)
	if target == nil {
		t.Error("instance is nil")
	}
}
