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
	last  *SLLNode[T]
}

// Size returns size of the list
func (s *SinglyLinkedList[T]) Size() int {
	return s.count
}

// GetFirstValue returns first value in the list
func (s *SinglyLinkedList[T]) GetFirstValue() (res T, err error) {

	return s.head.Value, nil
}

// Add add to the list
func (s *SinglyLinkedList[T]) Add(values ...T) {
	for _, value := range values {
		s.insertAt(s.Size(), value)
	}
}

// InsertAt insert value at specific index in the list
func (s *SinglyLinkedList[T]) insertAt(index int, value T) error {
	return nil
}
