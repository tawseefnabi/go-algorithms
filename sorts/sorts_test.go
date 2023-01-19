package sorts

import (
	"github.com/stretchr/testify/assert"
	// "github.com/tawseefnabi/go-algorithms/sorts"
	// "math/rand"
	"reflect"
	"testing"
	// "time"
	// "fmt"
)

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{10, 3, 2, 5, 4, 7, 6, 9, 8, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned",
		},
		//Random order with repetitions
		{
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
			name:     "Random order Signed",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			sorted := reflect.DeepEqual(actual, test.expected)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}

//BEGIN TESTS

func Test_Bubble(t *testing.T) {
	testFramework(t, Bubble[int])
}

func Test_Selection(t *testing.T) {
	testFramework(t, Selection[int])
}

func Test_BubbleSort(t *testing.T) {
	array := []int{43, 11, 18, 8, 3, 17}
	expectedBubbleSortArray := []int{3, 8, 11, 17, 18, 43}
	assert.Equal(t, BubbleSort(array), expectedBubbleSortArray)
}

func Test_SelectionSort(t *testing.T) {
	array := []int{43, 11, 18, 8, 3, 17}
	expectedSelectionSortArray := []int{3, 8, 11, 17, 18, 43}
	assert.Equal(t, SelectionSort(array), expectedSelectionSortArray)
}
