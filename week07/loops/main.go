package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// var length float64 = 3.2
	// var width int = 2
	// fmt.Println("면적은", int(length)*width)
	// fmt.Println("length > width?", int(length) > width)
	// fmt.Println("형변환", reflect.TypeOf(int(length)))
	// fmt.Println("원본", reflect.TypeOf(length))

	// var now time.Time = time.Now()
	// var month time.Month = now.Month() // month := now.Month()
	// var day int = now.Day()
	// fmt.Println(month, day)

	// // var univ string = "Go$ Inha$"
	// changer := strings.NewReplacer("$", "!")
	// // changed := changer.Replace(univ)
	// changed := changer.Replace("Go$ Inha$")
	// fmt.Println(changed)

	// r := bufio.NewReader(os.Stdin)
	// // i, _ := r.ReadString('\n') //ignore error
	// i, err := r.ReadString('\n')
	// // fmt.Println(err)
	// if err != nil {
	// 	log.Fatal(err) // report the error and exit the program
	// }
	// fmt.Println(i)

	// //shadowing
	// var float64 float64 = 2.71
	// var f float64 = 3.991
	// fmt.Println(float64)
	// fmt.Println(f)

	// var fmt float64 = 2.71
	// fmt.Println(fmt)

	// fmt.Print("Enter a score: ")
	// r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공란 및 탭 키 등 제거
	// score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수타입으로 변환
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var status string
	// if score >= 60 {
	// 	status = "Pass"
	// } else {
	// 	status = "Fail"
	// }
	// fmt.Println(score, status)

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
