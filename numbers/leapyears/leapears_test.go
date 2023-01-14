package leapyears

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsLeapYear(t *testing.T) {
	for year := 2004; year <= 2096; year = year + 4 {
		assert.False(t, IsLeapYear(year))
	}
	assert.False(t, IsLeapYear(1900))
	assert.True(t, IsLeapYear(2022))
	assert.False(t, IsLeapYear(2004))
}
