package fibonaccis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Fibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(0), 0)
	assert.Equal(t, Fibonacci(1), 1)
	assert.Equal(t, Fibonacci(2), 1)
	assert.Equal(t, Fibonacci(3), 2)
	assert.Equal(t, Fibonacci(6), 8)
	assert.Equal(t, Fibonacci(10), 55)
}
