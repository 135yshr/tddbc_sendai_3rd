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

func Test正常_開区間を作成し上端点が取得できることを確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		expected int
	}{
		{3, 8, 8},
		{3, 9, 9},
		{3, 7, 7},
	}

	for _, d := range data {
		expected := d.expected
		target := NewOpenRange(d.lower, d.upper)
		if target.upper != expected {
			t.Errorf("failed get lower expected=%d, actual=%d\n", expected, target.lower)
		}
	}
}

func Test正常_開区間の文字列表記が取得できることを確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		expected string
	}{
		{3, 8, "[3,8]"},
		{3, 9, "[3,9]"},
	}

	for _, d := range data {
		expected := d.expected
		target := NewOpenRange(d.lower, d.upper)
		actual := target.String()
		if actual != expected {
			t.Errorf("failed get lower expected=%s, actual=%s\n", expected, actual)
		}
	}
}

func Test正常_開区間に指定した整数を含むか確認する(t *testing.T) {
	var data = []struct {
		lower    int
		upper    int
		value    int
		expected bool
	}{
		{3, 8, 4, true},
		{3, 8, 7, true},
		{3, 8, 2, false},
		{3, 8, 9, false},
		{3, 8, 3, false},
		{3, 8, 8, false},
	}

	for _, d := range data {
		expected := d.expected
		value := d.value
		target := NewOpenRange(d.lower, d.upper)
		actual := target.Contains(value)
		if actual != expected {
			t.Errorf("failed %s value=%d expected=%t actual=%t\n", target, value, expected, actual)
		}
	}
}

func Test正常_開区間が別の開区間と等しいか確認する(t *testing.T) {
	var data = []struct {
		target   OpenRange
		value    OpenRange
		expected bool
	}{
		{OpenRange{3, 8}, OpenRange{3, 8}, true},
		{OpenRange{3, 8}, OpenRange{4, 8}, false},
		{OpenRange{3, 8}, OpenRange{3, 9}, false},
		{OpenRange{3, 8}, OpenRange{2, 8}, false},
		{OpenRange{3, 8}, OpenRange{3, 7}, false},
		{OpenRange{3, 8}, OpenRange{2, 9}, false},
	}
	for _, d := range data {
		expected := d.expected
		value := d.value
		target := d.target
		actual := target.Equal(value)
		if actual != expected {
			t.Errorf("failed %v value=%v expected=%t actual=%t\n", target, value, expected, actual)
		}
	}
}
