package main

import "fmt"

func main() {
	subjects := []string{"Go", "", "Python"}
	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
