package powers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowers(t *testing.T) {
	assert.Equal(t, Power(2, -3), 1.0/8)
	assert.Equal(t, Power(2, -2), 1.0/4)
	assert.Equal(t, Power(2, -1), 1.0/2)
	assert.Equal(t, Power(2, 0), 1.0)
	assert.Equal(t, Power(2, 1), 2.0)
	assert.Equal(t, Power(2, 2), 4.0)
	assert.Equal(t, Power(2, 3), 8.0)
}
