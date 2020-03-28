package goscore

// Each - Performs action on each element
func (l List) Each(executor func(i interface{})) {
	for _, elem := range l {
		executor(elem)
	}
}
