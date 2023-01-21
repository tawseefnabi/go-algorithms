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

type DynamicArray struct {
	Size        int
	Capacity    int
	ElementData []any
}
