package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Sort_Ascending(t *testing.T) {
	l := List{23, 12, 2, 4}
	s := l.Sort(nil)
	assert.Equal(t, s[0], 2)
	assert.Equal(t, s[1], 4)
	assert.Equal(t, s[2], 12)
	assert.Equal(t, s[3], 23)
}

func TestList_Sort_Descending(t *testing.T) {
	l := List{23, 12, 2, 4}
	s := l.Sort(func(i, j int) bool {
		return l[i].(int) > l[j].(int)
	})
	assert.Equal(t, s[3], 2)
	assert.Equal(t, s[2], 4)
	assert.Equal(t, s[1], 12)
	assert.Equal(t, s[0], 23)
}
