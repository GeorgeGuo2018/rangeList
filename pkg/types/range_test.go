package types

import (
	"testing"
)

func TestNewRange(t *testing.T) {
	_, err := NewRange(2, 3)
	if err != nil {
		t.Errorf("NewRange(2, 3) failed, expected err nil, got %s", err)
	}

	_, err = NewRange(2, 2)
	if err == nil {
		t.Errorf("NewRange(2, 2) failed, expected err, got nil")
	}

	_, err = NewRange(2, 1)
	if err == nil {
		t.Errorf("NewRange(2, 1) failed, expected err, got nil")
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
