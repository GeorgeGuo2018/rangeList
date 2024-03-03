package types

import "fmt"

type Range struct {
	begin int
	end   int
}

func NewRange(begin, end int) (Range, error) {
	if begin >= end {
		return Range{}, fmt.Errorf("wrong param, correct usage: NewRange(begin,end int) with begin less than end")
	}
	return Range{begin, end}, nil
}

func (arange Range) IsEqual(rg Range) bool {
	if arange.begin == rg.begin && arange.end == rg.end {
		return true
	}
	return false
}

func (arange Range) Value() []int {
	var value []int

	for i := arange.begin; i < arange.end; i++ {
		value = append(value, i)
	}
	return value
}
