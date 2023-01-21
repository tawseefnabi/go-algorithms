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

	// Delete: deletes the passed element from the set if present
	Delete(item any)

	// GetItems: gives the array( []any ) of elements of the set.
	GetItems() []any

	In(item any) bool
	// IsSubsetOf: checks whether set is subset of set2 or not
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
func (st *set) Delete(value any) {
	delete(st.elements, value)
}
func (st *set) GetItems() []any {
	keys := make([]any, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set) In(value any) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}
