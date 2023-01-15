package countdigits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Countdigits(t *testing.T) {
	assert.Equal(t, Countdigits(0), 1)
	assert.Equal(t, Countdigits(1), 1)
	assert.Equal(t, Countdigits(100), 3)
	assert.Equal(t, Countdigits(1012), 4)
}
