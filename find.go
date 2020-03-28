package goscore

func (l List) Find(i interface{}) ([]int, bool) {
	flag := false
	var indices []int
	for index, elem := range l {
		if elem == i {
			flag = true
			indices = append(indices, index)
		}
	}
	return indices, flag
}
