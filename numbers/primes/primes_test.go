package primes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimes(t *testing.T) {
	fmt.Println("Testing Reverse of Number: ")
	assert.True(t, IsPrime(2))
	assert.False(t, IsPrime(-1))
	assert.True(t, IsPrime(7))
	assert.False(t, IsPrime(12))
}
