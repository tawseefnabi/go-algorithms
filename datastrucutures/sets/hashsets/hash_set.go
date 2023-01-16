package hashsets

// New factory that creates a hash set
func New[T comparable](values ...T) *HashSet[T] {
	set := HashSet[T]{data: make(map[T]struct{}, len(values))}
	set.Add(values...)
	return &set
}

// HashSet datastructure
type HashSet[T comparable] struct {
	data map[T]struct{}
}

// Add adds values to the set
func (s *HashSet[T]) Add(values ...T) {
	for _, value := range values {
		s.data[value] = struct{}{}
	}
}

// IsEmpty checks if the set is empty

func (s *HashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the set

func (s *HashSet[T]) Size() int {
	return len(s.data)
}

// Contains checks if the value is in the set
func (s *HashSet[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists

}

// ContainsAll checks if all values are in the set
func (s *HashSet[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

// ContainsAny checks if any of the values are in the set
func (s *HashSet[T]) ContainsAny(values ...T) bool {
	for _, value := range values {
		if s.Contains(value) {
			return true
		}
	}
	return false
}

// Clear clears set
func (s *HashSet[T]) Clear() {
	s.data = make(map[T]struct{})
}

// GetValues returns values
func (s *HashSet[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	for key := range s.data {
		values = append(values, key)
	}
	return values
}
