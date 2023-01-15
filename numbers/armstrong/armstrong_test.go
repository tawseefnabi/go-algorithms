package armstrong

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsArmstrong(t *testing.T) {
	assert.True(t, IsArmstrong(0))
	assert.True(t, IsArmstrong(1))
	assert.True(t, IsArmstrong(153))
	assert.True(t, IsArmstrong(370))
	assert.False(t, IsArmstrong(120))
}
