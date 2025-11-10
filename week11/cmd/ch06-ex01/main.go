package main

import "fmt"

func main() {
	subjects := []string{"Go", "javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3] // slicing
	subjects[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go") // python : subjectsSlice.append("Go")
	//subjectsSlice = append(subjectsSlice, "Go", "DB") // 배열의 길이를 오버하면 다른 메모리로 할당
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("=======================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
