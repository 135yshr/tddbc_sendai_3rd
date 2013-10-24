package tddbc

import (
	"testing"
)

func Test正常_開区間を作成し下端点が取得できることを確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		expected int
	}{
		{3, 8, 3},
		{1, 8, 1},
		{2, 8, 2},
	}

	for _, d := range data {
		expected := d.expected
		target := NewOpenRange(d.lower, d.upper)
		if target.lower != expected {
			t.Errorf("failed get lower expected=%d, actual=%d\n", expected, target.lower)
		}
	}
}
