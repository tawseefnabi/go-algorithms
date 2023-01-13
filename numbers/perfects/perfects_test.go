package perfects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsPerfect(t *testing.T) {
	assert.False(t, IsPerfect(0))
	assert.False(t, IsPerfect(1))
	assert.False(t, IsPerfect(2))
	assert.False(t, IsPerfect(3))
	assert.False(t, IsPerfect(4))
	assert.False(t, IsPerfect(5))
	assert.True(t, IsPerfect(6))
	assert.True(t, IsPerfect(28))
	assert.True(t, IsPerfect(496))
	assert.True(t, IsPerfect(8128))
}
