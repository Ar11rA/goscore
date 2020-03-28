package goscore

import (
	"sort"
)

func (l List) Sort(sorter func(param1 int, param2 int) bool) List {
	clone := append(l[:0:0], l...)
	if sorter == nil {
		sorter = func(i int, j int) bool {
			return clone[i].(int) < clone[j].(int)
		}
	}
	sort.Slice(clone, sorter)
	return clone
}
