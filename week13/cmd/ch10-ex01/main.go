package main

import (
	"fmt"
	"log"
	calendar "week13/pkg/calender"
)

func main() {
	today := calendar.Event{} // Date가 임베딩되어 있음
	//err := today.SetTitle("Final Exam D-14.................................")
	err := today.SetTitle("Final Exam D-14")
	if err != nil {
		log.Fatal(err)
	}

	//today.year = 2025
	err = today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	//err = today.SetDay(77)
	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(today.Year(), "년", today.Month(), "월", today.Day(), "일")
	fmt.Printf("%s\n%d년 %d월 %d일\n", today.Title(), today.Year(), today.Month(), today.Day())
}
