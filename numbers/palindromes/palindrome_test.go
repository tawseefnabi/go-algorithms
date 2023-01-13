package palindromes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(12321))
	assert.True(t, IsPalindrome(8998))
	assert.False(t, IsPalindrome(12345))
}
