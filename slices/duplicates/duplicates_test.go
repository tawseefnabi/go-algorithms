package duplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ContainsDuplicate(t *testing.T) {
	assert.Equal(t, ContainsDuplicates(1, 2, 3), false)
	assert.Equal(t, ContainsDuplicates(1, 2, 3, 3), true)
	assert.Equal(t, ContainsDuplicates(true, false), false)
	assert.Equal(t, ContainsDuplicates(true, false, true), true)
	assert.Equal(t, ContainsDuplicates("1", "2", "3"), false)
	assert.Equal(t, ContainsDuplicates("1", "2", "1"), true)
}
