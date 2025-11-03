package main

import (
	"fmt"
)

func main() {
	// var arrayBool [3]bool = [3]bool{true, false, true}
	// arrayBool := [3]bool{true, false, true}
	// var arrayInt [3]int
	// fmt.Println(reflect.TypeOf(arrayBool)) // [3]bool
	// fmt.Printf("%#v\n", arrayBool)         // [3]bool{true, false, true}

	// fmt.Println(arrayBool[1])
	// arrayInt[1] = 2
	// fmt.Println(arrayInt[1])

	arrayBool := [3]bool{true, false, true}
	arrayInt := [3]int{-9, 11, 7}
	for i := 0; i < 3; i++ {
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayInt[i])
	}

}
