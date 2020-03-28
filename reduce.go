package goscore

// Accumulates result from given function and returns the result
func (l List) Reduce(reducer func(acc interface{}, i interface{}) interface{}, acc interface{}) interface{} {
	for _, elem := range l {
		acc = reducer(acc, elem)
	}
	return acc
}
