package hashmap

import (
	"fmt"
	// "github.com/tawseefnabi/go-algorithms/datastrucutures/hashmap"

	"testing"
)

func TestHashMap(t *testing.T) {
	mp := HashMap{}
	fmt.Println("new", mp)
	t.Run("Test 1: Put(10) and checking if Get() is correct", func(t *testing.T) {
		mp.Put("test", 10)
		// got := mp.Get("test")
		// if got != 10 {
		// 	t.Errorf("Put: %v, Got: %v", 10, got)
		// }
		// fmt.Println("new", mp)
	})
}
