package singlylinkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
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
func TestSinglyLinked_GetLastValue(t *testing.T) {
	list := New[string]()
	_, err := list.GetLastValue()
	assert.NotNil(t, err)

	list.Add("abc")
	value, err := list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "abc")
	list.Add("def")
	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "def")

	list.insertAt(2, "xyz")
	value, err = list.GetLastValue()
	assert.Nil(t, err)
	assert.Equal(t, value, "xyz")

}

func TestSinglyLinked_GetValues(t *testing.T) {
	list := New[string]()
	assert.True(t, reflect.DeepEqual(list.GetValues(), []string{}))

	list.Add("a", "b", "c")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []string{"a", "b", "c"}))

	list.insertAt(1, "d")
	assert.True(t, reflect.DeepEqual(list.GetValues(), []string{"a", "d", "b", "c"}))

}
