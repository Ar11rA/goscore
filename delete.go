package goscore

// Deletes element at particular index
func (l List) DeleteIndex(i int) List {
	if i > 0 && i < len(l) {
		updatedList := append(l[:i], l[i+1:]...)
		return updatedList
	}
	return l
}

// Deletes all elements with the passed value
func (l List) DeleteAllInstances(n int) List {
	var updatedList List
	for _, val := range l {
		if val != n {
			updatedList = append(updatedList, val)
		}
	}
	return updatedList
}

// Deletes first instance of element with the passed value
func (l List) DeleteFirstInstance(n int) List {
	indexToDelete := -1
	for index, val := range l {
		if val == n {
			indexToDelete = index
			break
		}
	}
	if indexToDelete == -1 {
		return l
	}
	updatedList := l.DeleteIndex(indexToDelete)
	return updatedList
}
