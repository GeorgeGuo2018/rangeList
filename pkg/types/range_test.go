package types

import (
	"fmt"
	"testing"
)

func TestNewRange(t *testing.T) {
	_, err := NewRange(2, 3)
	if err != nil {
		fmt.Print("2")
		t.Errorf("NewRange(2, 3) failed, expected err nil, got %s", err)
	}
	fmt.Print("3")
	_, err = NewRange(2, 2)
	if err == nil {
		t.Errorf("NewRange(2, 2) failed, expected err, got nil")
	}

	_, err = NewRange(2, 1)
	if err == nil {
		t.Errorf("NewRange(2, 1) failed, expected err, got nil")
	}
}

func TestEqual(t *testing.T) {
	testRange := Range{0, 3}
	testCases := []struct {
		Range
		expected bool
	}{
		{Range{0, 2}, false},
		{Range{-1, 1}, false},
		{Range{0, 3}, true},
		{Range{-2, 3}, false},
	}

	for _, tc := range testCases {
		if tc.Range.IsEqual(testRange) != tc.expected {
			t.Errorf("%+vEqual(%+v) = %t; expected %t", tc.Range, testRange, tc.Range.IsEqual(testRange), tc.expected)
		}
	}
}

func TestValue(t *testing.T) {
	testCases := []struct {
		Range
		expected []int
	}{
		{Range{1, 2}, []int{1}},
		{Range{-1, 1}, []int{-1, 0}},
		{Range{0, 3}, []int{0, 1, 2}},
		{Range{-2, 0}, []int{-2, -1}},
	}

	for _, tc := range testCases {
		result := tc.Range.Value()
		for idx, value := range result {
			if value != tc.expected[idx] {
				t.Errorf("Value(%d, %d) = %+v; expected %+v", tc.Range.begin, tc.Range.end, result, tc.expected)
			}
		}
	}
}
