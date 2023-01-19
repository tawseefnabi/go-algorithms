package set

func New(items ...any) Set {
	st := set{
		elements: make(map[any]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

// Set is an interface of possible methods on 'set'.
type Set interface {

	// Len: gives the length of the set (total no. of elements in set)
	Len() int

	// Add: adds new element to the set
	Add(item any)
}

type set struct {
	elements map[any]bool
}

func (st *set) Add(value any) {
	st.elements[value] = true
}
func (st *set) Len() int {
	return len(st.elements)
}
