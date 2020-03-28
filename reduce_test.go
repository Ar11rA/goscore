package goscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func adder(acc interface{}, i interface{}) interface{} {
	return acc.(int) + i.(int)
}

func frequency(acc interface{}, i interface{}) interface{} {
	if _, ok := acc.(map[int]int)[i.(int)]; ok {
		acc.(map[int]int)[i.(int)] = acc.(map[int]int)[i.(int)] + 1
	} else {
		acc.(map[int]int)[i.(int)] = 1
	}
	return acc
}

func TestList_Reduce_WithSum(t *testing.T) {
	l := List{1, 2, 3, 4, 5}
	sum := l.Reduce(adder, 0)
	assert.Equal(t, sum, 15)
}

func TestList_Reduce_WithGroup(t *testing.T) {
	l := List{1, 2, 1, 1, 2}
	group := l.Reduce(frequency, map[int]int{})
	assert.Equal(t, group.(map[int]int)[1], 3)
	assert.Equal(t, group.(map[int]int)[2], 2)
}
