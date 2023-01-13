package reverses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse(12321), 12321)
	assert.NotEqual(t, Reverse(321), 345)
	assert.Equal(t, Reverse(8998), 8998)
	assert.Equal(t, Reverse(12345), 54321)
}
