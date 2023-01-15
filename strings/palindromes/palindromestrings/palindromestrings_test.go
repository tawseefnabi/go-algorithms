package palindromestrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("hannah"))
	assert.True(t, IsPalindrome("hanah"))
	assert.False(t, IsPalindrome("hello"))
}
