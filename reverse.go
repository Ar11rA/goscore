package goscore

func (l List) Reverse() List {
	s := make(List, len(l))
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = l[j], l[i]
	}
	return s
}