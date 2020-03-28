package goscore

// Unique - Returns only unique set of elements
func (l List) Unique() List {
	keys := make(map[interface{}]bool)
	var list List
	for _, entry := range l {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
