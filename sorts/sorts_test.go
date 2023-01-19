package sorts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	array := []int{43, 11, 18, 8, 3, 17}
	bubbleSortArray := []int{3, 8, 11, 17, 18, 43}
	assert.Equal(t, BubbleSort(array), bubbleSortArray)
}

func Test_SelectionSort(t *testing.T) {
	array := []int{43, 11, 18, 8, 3, 17}
	SelectionSortArray := []int{3, 8, 11, 17, 18, 43}
	assert.Equal(t, SelectionSort(array), SelectionSortArray)
}
