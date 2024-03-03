package types

import (
	"strconv"
)

type RangeList struct {
	elems []Range
}

func NewRangeList() *RangeList {
	rangeList := make([]Range, 0)
	return &RangeList{elems: rangeList}
}

func (rangeList *RangeList) Add(arange Range) *RangeList {
	if rangeList == nil {
		rangeList = NewRangeList()
	}
	if rangeList.elems == nil {
		rangeList.elems = make([]Range, 0)
	}
	rangeList.elems = append(rangeList.elems, arange)
	return rangeList
}

func (rangeList *RangeList) Remove(arange Range) *RangeList {
	if rangeList == nil {
		return NewRangeList()
	}
	if rangeList.elems == nil {
		rangeList.elems = make([]Range, 0)
		return rangeList
	}

	for idx, elem := range rangeList.elems {
		if arange.IsEqual(elem) {
			if len(rangeList.elems) == 1 {
				rangeList.elems = make([]Range, 0)
				break
			}
			if idx == 0 {
				rangeList.elems = rangeList.elems[1:]
				break
			}
			if idx == len(rangeList.elems)-1 {
				rangeList.elems = rangeList.elems[:len(rangeList.elems)-1]
				break
			}
		}
	}

	return rangeList
}

func (rangeList *RangeList) ToString() string {
	result := ""
	if rangeList == nil || len(rangeList.elems) == 0 {
		return result
	}
	for _, elem := range rangeList.elems {
		result = result + " [" + strconv.Itoa(elem.begin) + "," + strconv.Itoa(elem.end) + ")"
	}

	//remove the beginning space
	result = result[1:]
	return result
}
