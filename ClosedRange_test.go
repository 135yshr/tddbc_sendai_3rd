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

func Test正常_閉区間3to8を作成し下端点が取得できることを確認する(t *testing.T) {
	expected := 3
	target := NewClosedRange(3, 8)
	if target.lower != expected {
		t.Errorf("failed get lower expected=%d, actual=%d\n", expected, target.lower)
	}
}
