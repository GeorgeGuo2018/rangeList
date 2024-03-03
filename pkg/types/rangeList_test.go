package types

import (
	"testing"
)

func TestNewRangeList(t *testing.T) {
	testCases := []struct {
		testCase []Range
		expected RangeList
	}{

		{nil, RangeList{[]Range{}}},
		{[]Range{}, RangeList{[]Range{}}},
		{[]Range{Range{0, 3}}, RangeList{[]Range{Range{0, 3}}}},
		{[]Range{Range{-2, 3}, Range{-2, 3}}, RangeList{[]Range{Range{-2, 3}, {-2, 3}}}},
		{[]Range{Range{0, 3}, Range{-2, 3}}, RangeList{[]Range{Range{0, 3}, {-2, 3}}}},
	}

	for _, tc := range testCases {
		if !NewRangeList(tc.testCase).IsEqual(&tc.expected) {
			t.Errorf("NewRangeList with %+v get %+v; expected %+v", tc.testCase, NewRangeList(tc.testCase), tc.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	testRangeList := NewRangeList(nil)
	testCases := [5]struct {
		Range
		expected RangeList
	}{
		{Range{0, 2}, RangeList{[]Range{Range{0, 2}}}},
		{Range{-1, 1}, RangeList{[]Range{Range{-1, 1}, {0, 2}}}},
		{Range{0, 3}, RangeList{[]Range{Range{0, 3}, {-1, 1}, {0, 2}}}},
		{Range{-2, 3}, RangeList{[]Range{Range{-2, 3}, {0, 3}, {-1, 1}, {0, 2}}}},
		{Range{-2, 3}, RangeList{[]Range{Range{-2, 3}, {-2, 3}, {0, 3}, {-1, 1}, {0, 2}}}},
	}

	for _, tc := range testCases {
		if !testRangeList.Add(tc.Range).IsEqual(&tc.expected) {
			t.Errorf("after add %+v get %+v; expected %+v", tc.Range, testRangeList, tc.expected)
		}
	}
}

func TestRemove(t *testing.T) {
	testRangeList := NewRangeList([]Range{Range{-2, 3}, {-2, 3}, {0, 3}, {-1, 1}, {0, 2}})
	testCases := [5]struct {
		Range
		expected RangeList
	}{
		{Range{0, 2}, RangeList{[]Range{{-2, 3}, {-2, 3}, {0, 3}, {-1, 1}}}},
		{Range{-1, 1}, RangeList{[]Range{{-2, 3}, {-2, 3}, {0, 3}}}},
		{Range{0, 3}, RangeList{[]Range{{-2, 3}, {-2, 3}}}},
		{Range{-2, 3}, RangeList{[]Range{{-2, 3}}}},
		{Range{-2, 3}, RangeList{[]Range{}}},
	}

	for _, tc := range testCases {
		if !testRangeList.Remove(tc.Range).IsEqual(&tc.expected) {
			t.Errorf("after remove %+v get %+v; expected %+v", tc.Range, testRangeList, tc.expected)
		}
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		testCase *RangeList
		expected string
	}{

		{nil, ""},
		{&RangeList{[]Range{}}, ""},
		{&RangeList{[]Range{Range{0, 3}}}, "[0,3)"},
		{&RangeList{[]Range{Range{-2, 3}, Range{-2, 3}}}, "[-2,3) [-2,3)"},
		{&RangeList{[]Range{Range{0, 3}, Range{-2, 3}}}, "[0,3) [-2,3)"},
	}

	for _, tc := range testCases {
		if tc.testCase.ToString() != tc.expected {
			t.Errorf("ToString with %+v get %+v; expected %+v", tc.testCase, tc.testCase.ToString(), tc.expected)
		}
	}
}

func TestIsEqual(t *testing.T) {
	testRangeList := RangeList{[]Range{{0, 3}, {0, 3}}}
	testCases := []struct {
		RangeList
		expected bool
	}{
		{RangeList{[]Range{{0, 2}}}, false},
		{RangeList{[]Range{{0, 2}, {0, 3}}}, false},
		{RangeList{[]Range{{0, 3}, {0, 3}}}, true},
	}

	for _, tc := range testCases {
		if testRangeList.IsEqual(&tc.RangeList) != tc.expected {
			t.Errorf("%+vEqual(%+v) = %t; expected %t", testRangeList, tc.RangeList, testRangeList.IsEqual(&tc.RangeList), tc.expected)
		}
	}
}
