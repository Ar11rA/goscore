package goscore

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func double(item interface{}) interface{} {
	return item.(int) * 2
}

func upperCase(item interface{}) interface{} {
	return strings.ToUpper(item.(string))
}

func addZ(item interface{}) interface{} {
	return strconv.Itoa(item.(int)) + "Z"
}

func TestList_Map_WithDouble(t *testing.T) {
	l := List{1, 2}
	mapped := l.Map(double)
	assert.Equal(t, mapped[0], 2)
	assert.Equal(t, mapped[1], 4)
}

func TestList_Map_WithUpperCase(t *testing.T) {
    l := List{"abc", "def"}
	mapped := l.Map(upperCase)
	assert.Equal(t, mapped[0], "ABC")
	assert.Equal(t, mapped[1], "DEF")
}

func TestList_Map_WithDifferentTypes(t *testing.T) {
	l := List{1, 2}
	mapped := l.Map(addZ)
	assert.Equal(t, mapped[0], "1Z")
	assert.Equal(t, mapped[1], "2Z")
}
