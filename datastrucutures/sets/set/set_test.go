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
