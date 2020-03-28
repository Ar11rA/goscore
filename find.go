package goscore

// Find - Finds the passed element and return indices of occurrence and returns true if present
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
