package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println(math.Round(2.21)) //반올림
	// fmt.Println(math.Ceil(2.21))  //올림
	// fmt.Println(math.Floor(2.21)) //내림
	// fmt.Println(strings.Title("go developer!"))
	// //각 단어의 첫글자를 대문자로 바꿔서 출력(strings.Title은 현재 '비권장'처리, 유니코드 처리가 완벽하지 않음)
	// fmt.Println("Kim\nInha\t\"\\") //"\n" : 줄바꿈 / "\t" : tap / 역슬래시 다음에 있는 특수문자 ",\를 출력
	// fmt.Println('2', '가')          //''안에 있는 문자는 rune(int32)로 처리, 유니코드 값으로 출력 '2' = "50" / '가' = "44032"

	// fmt.Println()
	// fmt.Println(reflect.TypeOf(2.31))       //float64
	// fmt.Println(reflect.TypeOf("Kim inha")) //string
	// fmt.Println(reflect.TypeOf(true))       //bool
	// fmt.Println(reflect.TypeOf('A'))        //int32
	// fmt.Println(reflect.TypeOf(19))         //int

	// var id int16 //var 변수명 타입
	// var name string
	// var gpa float32
	// id = 999 //값 대입
	// name = "kimsangwon"
	// gpa = 3.99

	// var id int16 = 999	//명시적 선언
	// var name string = "kimsangwon"
	// var gpa float32 = 3.99

	// id := 999	//짧은 선언 - 타입 자동 추론
	// name := "kimsangwon"
	// gpa := 3.99

	// fmt.Println("학번은", id, reflect.TypeOf(id), ", 이름은", name, reflect.TypeOf(name))
	// //학번은 999 int16 , 이름은 kimsangwon string
	// fmt.Println("평점 : ", gpa, reflect.TypeOf(gpa))
	// //평점 :  3.99 float32

	// var f64 float64 //기본값만 적용
	// var t bool
	// var s string
	// var i int
	// var i16 int16

	// fmt.Println(f64, reflect.TypeOf(f64)) // 0 float64
	// fmt.Println(t, reflect.TypeOf(t))     // false bool
	// fmt.Println(s, reflect.TypeOf(s))     // (공백) string
	// fmt.Println(i, reflect.TypeOf(i))     // 0 int
	// fmt.Println(i16, reflect.TypeOf(i16)) // 0 int16

	var f64 float64 //변수 선언만
	// total_price := 1000 //스네이크 케이스를 사용한 이름(_로 구분)
	// totalprice := 1000	//일반적인 이름, 가독성이 떨어짐, GO언어는 대소문자 구분
	totalPrice := 1000 //camelCase를 사용한 이름(대문자로 구분)
	// 기본적으로 GO 언어는 camelCase 기반임

	fmt.Println(totalPrice)                      // 1000
	fmt.Println(totalPrice, reflect.TypeOf(f64)) // 1000 float64
}
