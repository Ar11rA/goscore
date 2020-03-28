package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Find_Sucess(t *testing.T) {
	l := List{1, 2, 3, 1, 2}
	indices, isPresent := l.Find(2)
	assert.Equal(t, isPresent, true)
	assert.Equal(t, len(indices), 2)
	assert.Equal(t, indices[0], 1)
	assert.Equal(t, indices[1], 4)
}

func TestList_Find_Fail(t *testing.T) {
	l := List{1, 2, 3, 1, 2}
	indices, isPresent := l.Find(4)
	assert.Equal(t, isPresent, false)
	assert.Equal(t, len(indices), 0)
}
