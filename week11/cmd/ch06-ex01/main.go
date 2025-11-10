package main

import "fmt"

func main() {
	subjects := []string{"Go", "javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3] // slicing
	//subjects[0] = "Java"
	subjectsSlice[0] = "Database"
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("=======================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
