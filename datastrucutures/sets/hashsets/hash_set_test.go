package hashsets

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HashSets(t *testing.T) {
	set := New[string]()
	assert.Equal(t, set.IsEmpty(), true)
	set.Add("hello")
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.Contains("hello"), true)
	assert.Equal(t, set.Contains("h"), false)
	assert.Equal(t, set.ContainsAll("hello", "unknown"), false)
	assert.Equal(t, set.ContainsAny("hello", "unknown"), true)
	assert.Equal(t, set.IsEmpty(), false)
	assert.Equal(t, set.Size(), 1)
	assert.Equal(t, set.GetValues()[0], "hello")
	set.Clear()
	assert.Equal(t, set.IsEmpty(), true)
	assert.Equal(t, set.Size(), 0)

}
