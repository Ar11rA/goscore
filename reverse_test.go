package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Reverse(t *testing.T) {
	l := List{1, 2, 3, 4}
	s := l.Reverse()
	assert.Equal(t, s[0], 4)
	assert.Equal(t, s[1], 3)
	assert.Equal(t, s[2], 2)
	assert.Equal(t, s[3], 1)
}
