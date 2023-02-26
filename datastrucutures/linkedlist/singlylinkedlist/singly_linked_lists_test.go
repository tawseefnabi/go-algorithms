package singlylinkedlist

import (
	// "fmt"
	// "reflect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglyLinked_GetFirstValue(t *testing.T) {
	list := New[string]()
	_, err := list.GetFirstValue()
	assert.NotNil(t, err)
	list.Add("abc")
	value, err := list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")
	list.Add("def")
	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")
	list.insertAt(0, "xyz")
	value, err = list.GetFirstValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "xyz")
}
