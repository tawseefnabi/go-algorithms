package singlylinkedlist

import (
	"errors"
)

var (
	errEmptyList        = errors.New("list is empty")
	errIndexOutOfBounds = errors.New("index is out of bounds")
)

// New factory to generate new singly linked lists
func New[T any](values ...T) *SinglyLinkedList[T] {
	list := SinglyLinkedList[T]{}
	list.Add(values...)
	return &list

}

// SLLNode singly linked list node
type SLLNode[T any] struct {
	Value T
	Next  *SLLNode[T]
}

// SinglyLinkedList singly linked list structure
type SinglyLinkedList[T any] struct {
	count int
	head  *SLLNode[T]
	tail  *SLLNode[T]
}

// IsEmpty checks if the list is empty
func (s *SinglyLinkedList[T]) isEmpty() bool {
	return s.Size() == 0
}

// Size returns size of the list
func (s *SinglyLinkedList[T]) Size() int {
	return s.count
}

// GetFirstValue returns first value in the list
func (s *SinglyLinkedList[T]) GetFirstValue() (res T, err error) {
	if s.isEmpty() {
		return res, errEmptyList
	}
	return s.head.Value, nil
}

// GetLastValue returns first value in the list
func (s *SinglyLinkedList[T]) GetLastValue() (res T, err error) {
	if s.isEmpty() {
		return res, errEmptyList
	}
	return s.tail.Value, nil
}

// GetValues returns values
func (s *SinglyLinkedList[T]) GetValues() []T {
	values := make([]T, 0, s.Size())
	current := s.head

	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	current = nil
	return values
}

// Add add to the list
func (s *SinglyLinkedList[T]) Add(values ...T) {
	for _, value := range values {
		s.insertAt(s.Size(), value)
	}
}

// InsertAt insert value at specific index in the list
func (s *SinglyLinkedList[T]) insertAt(index int, value T) error {
	if index < 0 || index > s.count {
		return errIndexOutOfBounds
	}
	element := &SLLNode[T]{Value: value}
	if s.isEmpty() {
		s.head = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}
	if index == 0 {
		element.Next = s.head
		s.head = element
		s.count++
		element = nil
		return nil
	}
	if index == s.count {
		s.tail.Next = element
		s.tail = element
		s.count++
		element = nil
		return nil
	}
	current, err := s.getNode(index - 1)

	if err != nil {
		return err
	}
	element.Next = current.Next
	current.Next = element
	s.count++
	current = nil
	element = nil
	return nil
}

func (s *SinglyLinkedList[T]) getNode(index int) (*SLLNode[T], error) {
	if index < 0 || index >= s.count {
		return nil, errIndexOutOfBounds
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current, nil
}
