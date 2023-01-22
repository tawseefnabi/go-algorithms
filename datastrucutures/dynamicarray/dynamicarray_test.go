package dynamicarray

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestDynamicArray(t *testing.T) {
	numbers := DynamicArray{}
	// check numbers is empty or nut
	t.Run("Check Empty Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != true {
			t.Errorf("Expected be true but got %v", numbers.IsEmpty())
		}
	})
	numbers.Add(10)
	numbers.Add(20)
	numbers.Add(30)
	numbers.Add(40)
	numbers.Add(50)

	// check numbers added to our dynamic array
	t.Run("Add Element into Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != false {
			t.Errorf("Expected be false but got %v", numbers.IsEmpty())
		}
		var res []any
		res = append(res, 10)
		res = append(res, 20)
		res = append(res, 30)
		res = append(res, 40)
		res = append(res, 50)
		if !reflect.DeepEqual(numbers.GetData(), res) {
			t.Errorf("Expected  be [10 20 30 40 50] but got %v", numbers.GetData())
		}
	})

	// Remove an Element inside the dynamic array with the index of array
	t.Run("Remove Element into Dynamic Array", func(t *testing.T) {
		if numbers.IsEmpty() != false {
			t.Errorf("Expected be false but got %v", numbers.IsEmpty())
		}
		var res []any
		res = append(res, 10)
		res = append(res, 20)
		res = append(res, 40)
		res = append(res, 50)

		// remove the element by the index
		err := numbers.Remove(2)
		if err != nil {
			t.Errorf("Expected be [10, 30, 40, 50] but got an Error %v", err)
		}
		if !reflect.DeepEqual(numbers.GetData(), res) {
			t.Errorf("Expected  be [10 20 40 50] but got %v", numbers.GetData())
		}
	})
}
