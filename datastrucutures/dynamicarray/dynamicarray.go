// Package dynamicarray
// A dynamic array is quite similar to a regular array, but its Size is modifiable during program runtime,
// very similar to how a slice in Go works. The implementation is for educational purposes and explains
// how one might go about implementing their own version of slices.
//
// For more details check out those links below here:
// GeeksForGeeks article : https://www.geeksforgeeks.org/how-do-dynamic-arrays-work/
// Go blog: https://blog.golang.org/slices-intro
// Go blog: https://blog.golang.org/slices
// authors [Wesllhey Holanda](https://github.com/wesllhey), [Milad](https://github.com/miraddo)
// see dynamicarray.go, dynamicarray_test.go
package dynamicarray

import (
	"errors"
	"fmt"
)

var defaultCapacity = 10

type DynamicArray struct {
	Size        int
	Capacity    int
	ElementData []any
}

// IsEmpty function is check that the array has value or not
func (da *DynamicArray) IsEmpty() bool {
	return da.Size == 0
}

// NewCapacity function increase the Capacity
func (da *DynamicArray) NewCapacity() {
	if da.Capacity == 0 {
		da.Capacity = defaultCapacity
	} else {
		fmt.Println("++++++++++++ELSE capa", da)
		da.Capacity = da.Capacity << 1
	}
	newElementData := make([]any, da.Capacity)
	copy(newElementData, da.ElementData)
	da.ElementData = newElementData
}

// Add function is add new element to our array
func (da *DynamicArray) Add(el any) {
	if da.Size == da.Capacity {
		da.NewCapacity()
	}
	da.ElementData[da.Size] = el
	da.Size++
}

// Remove function is remove an element with the index
func (da *DynamicArray) Remove(index int) error {
	err := da.CheckRangeFromIndex(index)
	if err != nil {
		return err
	}
	copy(da.ElementData[index:], da.ElementData[index+1:da.Size])
	da.ElementData[da.Size-1] = nil
	da.Size--
	return err
}

// GetData function return all value of array
func (da *DynamicArray) GetData() []any {
	return da.ElementData[:da.Size]

}

// CheckRangeFromIndex function it will check the range from the index
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.Size || index < 0 {
		return errors.New("out of range")
	}
	return nil
}
