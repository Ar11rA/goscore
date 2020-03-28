package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Unique(t *testing.T) {
	l := List{1, 1, 1, 2, 2, 2, 3, 3}
	s := l.Unique()
	assert.Equal(t, len(s), 3)
	assert.Equal(t, s[0], 1)
	assert.Equal(t, s[1], 2)
	assert.Equal(t, s[2], 3)
}