package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_DeleteIndex_Success(t *testing.T) {
	l := List{1, 2, 3}
	s := l.DeleteIndex(1)
	assert.Equal(t, len(s), 2)
	assert.Equal(t, s[0], 1)
	assert.Equal(t, s[1], 3)
}

func TestList_DeleteIndex_Fail(t *testing.T) {
	l := List{1, 2, 3}
	s := l.DeleteIndex(-1)
	assert.Equal(t, len(s), 3)
}

func TestList_DeleteAllInstances(t *testing.T) {
	l := List{1, 2, 3, 2, 1, 2}
	s := l.DeleteAllInstances(2)
	assert.Equal(t, len(s), 3)
	assert.Equal(t, s[0], 1)
	assert.Equal(t, s[1], 3)
	assert.Equal(t, s[2], 1)
}

func TestList_DeleteFirstInstance_Success(t *testing.T) {
	l := List{1, 2, 3, 2, 1, 2}
	s := l.DeleteFirstInstance(2)
	assert.Equal(t, len(s), 5)
	assert.Equal(t, s[0], 1)
	assert.Equal(t, s[1], 3)
	assert.Equal(t, s[2], 2)
	assert.Equal(t, s[3], 1)
	assert.Equal(t, s[4], 2)
}

func TestList_DeleteFirstInstance_Fail(t *testing.T) {
	l := List{1, 2, 3, 2, 1, 2}
	s := l.DeleteFirstInstance(4)
	assert.Equal(t, len(s), 6)
}