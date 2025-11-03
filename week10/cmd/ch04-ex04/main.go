package main

import (
	"week10/pkg/greeting"
	"week10/pkg/greeting/deutsch"
	korean "week10/pkg/greeting/kor"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.GutenTag()
	korean.Anyung() // 대소문자 구분함(함수명이 소문자로 시작하는 경우 외부 파일에서 접근 불가)
}
