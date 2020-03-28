package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isEven(num interface{}) bool {
	return num.(int)%2 == 0
}

func TestList_Filter(t *testing.T) {
	l := List{1, 2, 3, 4, 5}
	filteredList := l.Filter(isEven)
	assert.Equal(t, filteredList[0], 2)
	assert.Equal(t, filteredList[1], 4)
}
