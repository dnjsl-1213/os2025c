package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println(math.Round(2.21))
	// fmt.Println(math.Ceil(2.21))
	// fmt.Println(strings.Title("go developer!"))
	// fmt.Println("Kim\nInha\t\"\\")
	// fmt.Println('2', '가')

	// fmt.Println(reflect.TypeOf(2.31))
	// fmt.Println(reflect.TypeOf("Kim inha"))
	// fmt.Println(reflect.TypeOf(true))
	// fmt.Println(reflect.TypeOf('A'))
	// fmt.Println(reflect.TypeOf(19))

	// var id int16
	// var name string
	// var gpa float32

	// id = 999
	// name = "kimsangwon"
	// gpa = 3.99

	// var id int16 = 999
	// var name string = "kimsangwon"
	// var gpa float32 = 3.99

	// id := 999
	// name := "kimsangwon"
	// gpa := 3.99

	// fmt.Println("학번은", id, reflect.TypeOf(id), ", 이름은", name, reflect.TypeOf(name))
	// fmt.Println("평점 : ", gpa, reflect.TypeOf(gpa))

	// var f64 float64
	// var t bool
	// var s string
	// var i int
	// var i16 int16

	// fmt.Println(f64, reflect.TypeOf(f64))
	// fmt.Println(t, reflect.TypeOf(t))
	// fmt.Println(s, reflect.TypeOf(s))
	// fmt.Println(i, reflect.TypeOf(i))
	// fmt.Println(i16, reflect.TypeOf(i16))

	var f64 float64
	// total_price := 1000
	// totalprice := 1000
	totalPrice := 1000

	fmt.Println(totalPrice)
	fmt.Println(totalPrice, reflect.TypeOf(f64))
}
