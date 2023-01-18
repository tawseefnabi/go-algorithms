// Package constraints has some useful generic type constraints defined which is very similar to
// [golang.org/x/exp/constraints](https://pkg.go.dev/golang.org/x/exp/constraints) package.
// We refrained from using that until it gets placed into the standard library - currently
// there are some questions regarding this package [ref](https://github.com/golang/go/issues/50792).
package constraints

// Signed is a generic type constraint for all signed integers.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a generic type constraint for all unsigned integers.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
