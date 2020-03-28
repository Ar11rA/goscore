package goscore

import (
	"testing"
)

func doSomething(i interface{}) {
	i = i.(int) + 1
}
func TestList_Each(t *testing.T) {
	l := List{1, 2, 3}
	l.Each(doSomething)
}
