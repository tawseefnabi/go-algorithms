package reverses

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println("Testing Reverse of Number: ")
	assert.Equal(t, Reverse(12321), 12321)
	assert.NotEqual(t, Reverse(321), 345)
	assert.Equal(t, Reverse(8998), 8998)
	assert.Equal(t, Reverse(12345), 54321)
}
