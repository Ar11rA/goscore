package goscore

// Filter - Returns a filtered list of elements
func (l List) Filter(filter func(i interface{}) bool) List {
	var filteredList List
	for _, elem := range l {
		isFiltered := filter(elem)
		if isFiltered {
			filteredList = append(filteredList, elem)
		}
	}
	return filteredList
}
