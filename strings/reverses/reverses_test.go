package reversestrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Reverse(t *testing.T) {
	assert.Equal(t, ReverseString("hello"), "olleh")
}
