package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var arrayBool [3]bool = [3]bool{true, false, true}
	arrayBool := [3]bool{true, false, true}
	var arrayInt [3]int
	fmt.Println(reflect.TypeOf(arrayBool)) // [3]bool
	fmt.Printf("%#v\n", arrayBool)         // [3]bool{true, false, true}

	fmt.Println(arrayBool[1])
	arrayInt[1] = 2
	fmt.Println(arrayInt[1])
}
