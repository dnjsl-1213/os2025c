package main

import (
	"fmt"
	calendar "week13/pkg/calender"
)

func main() {
	today := calendar.Date{}
	//today.year = 2025
	today.SetYear(2025)
	today.SetMonth(11)
	today.SetDay(24)
	//fmt.Println(today.Year(), "년", today.Month(), "월", today.Day(), "일")
	fmt.Printf("%d년 %d월 %d일\n", today.Year(), today.Month(), today.Day())
}
