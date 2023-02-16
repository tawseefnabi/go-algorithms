package singlylinkedlist

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinked_GetFirstValue(t *testing.T) {
	list := New[string]()
	fmt.Println("++++", list)
}
