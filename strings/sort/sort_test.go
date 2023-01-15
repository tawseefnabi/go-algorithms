package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sort(t *testing.T) {
	assert.Equal(t, Sort(""), "")
	assert.Equal(t, Sort("a"), "a")
	assert.Equal(t, Sort("cba"), "abc")

}
