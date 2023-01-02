package main

import (
	"fmt"
	"reflect"
)

// the function signature accepts an arbitrary number of arguments of type slice.

func variadicExample(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.TypeOf(v), reflect.ValueOf(v).Kind())
	}

}
func main() {
	variadicExample(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}
