package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var s = "Japan"

func main() {
	strVar := "100"

	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}
