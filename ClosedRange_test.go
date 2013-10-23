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

func Test正常_閉区間を作成し下端点が取得できることを確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		expected int
	}{
		{1, 8, 1},
		{2, 8, 2},
		{3, 8, 3},
		{4, 8, 4},
	}

	for _, d := range data {
		expected := d.expected
		target := NewClosedRange(d.lower, d.upper)
		if target.lower != expected {
			t.Errorf("failed get lower expected=%d, actual=%d\n", expected, target.lower)
		}
	}
}

func Test正常_閉区間を作成し上端点が取得できることを確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		expected int
	}{
		{3, 7, 7},
		{3, 8, 8},
		{3, 9, 9},
		{3, 10, 10},
	}
	for _, d := range data {
		expected := d.expected
		target := NewClosedRange(d.lower, d.upper)
		if target.upper != expected {
			t.Errorf("failed get lower expected=%d, actual=%d\n", expected, target.upper)
		}
	}
}

func Test異常_上端点が下端点よりも小さいときはインスタンスが作成されないことを確認する(t *testing.T) {
	target := NewClosedRange(8, 3)
	if target != nil {
		t.Error("An error did not occur.")
	}
}

func Test正常_閉区間が指定した整数を含むかどうかを確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		param    int
		expected bool
	}{
		{3, 5, 4, true},
		{3, 5, 3, true},
		{3, 5, 5, true},
	}

	for _, d := range data {
		expected := d.expected
		target := NewClosedRange(d.lower, d.upper)
		actual := target.Contains(d.param) == expected
		if actual == false {
			t.Errorf("failed contains test expected=%t, actual=%t\n", expected, actual)
		}
	}
}
