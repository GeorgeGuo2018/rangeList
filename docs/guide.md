# RangeList user guide

## Range
### 1. initial a range
```bash
//if end less than begin,this func would return a err
func NewRange(begin, end int) (Range, error)
```

### 2. judege whether two ranges are equal
```bash
//if two Range have the same begin and end,the two Range is equal
func (arange Range) IsEqual(rg Range) bool
```

### 3. get the value of a range in type of slice
```bash
func (arange Range) Value() []int
```

## RangeList
### 1. initial a range list
```bash
func NewRangeList(elems []Range) *RangeList
```

### 2. add a range into range list
```bash
func (rangeList *RangeList) Add(arange Range) *RangeList
```

### 3. remove a range from range list
```bash
func (rangeList *RangeList) Remove(arange Range) *RangeList 
```

### 4. convert a range list to string
```bash
func (rangeList *RangeList) ToString() string
```

### 5. judege whether two rangeLists are equal
```bash
//if two RangeList have the same Range elem set,the two RangeList is equal 
func (rangeList *RangeList) IsEqual(rgList *RangeList) bool
```