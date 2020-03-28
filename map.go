package goscore

// Runs given function for all elements and returns list
func (l List) Map(mapper func(i interface{}) interface{}) List {
	mapped := make(List, len(l))
	for index, elem := range l {
		mapped[index] = mapper(elem)
	}
	return mapped
}
