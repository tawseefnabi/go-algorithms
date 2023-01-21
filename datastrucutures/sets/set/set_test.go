package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {
	set := New(1, 2, 3)
	assert.Equal(t, set.Len(), 3)
	set.Add(2)
	assert.Equal(t, set.Len(), 3)
	set.Add(4)
	assert.Equal(t, set.Len(), 4)
}
func Test_Add(t *testing.T) {
	set := New(1, 2, 3)
	set.Add(5)
	assert.Equal(t, set.Len(), 4)
}

func Test_Delete(t *testing.T) {
	set := New(1, 2, 3)
	set.Delete(2)
	assert.Equal(t, set.Len(), 2)
}

func Test_In(t *testing.T) {
	set := New(1, 2, 3)
	assert.True(t, set.In(2))
}
