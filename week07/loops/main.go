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
	// fmt.Println("면적은", int(length)*width)               //면적은 6
	// fmt.Println("length > width?", int(length) > width) //length > width? true
	// fmt.Println("형변환", reflect.TypeOf(int(length)))     //형변환 int
	// fmt.Println("원본", reflect.TypeOf(length))           //원본 float64

	// var now time.Time = time.Now()     //현재 시간 정보 출력 (년-월-일 시:분:초.7숫자 +0900 KST m=+0.000000001) 24줄
	// var month time.Month = now.Month() //month := now.Month() / 25줄 / ex)October
	// var day int = now.Day()            //day := now.Day / 25줄 / ex)20
	// fmt.Println(now)
	// fmt.Println(month, day)

	// // var univ string = "Go$ Inha$"
	// changer := strings.NewReplacer("$", "!") //$ -> ! 변환
	// // changed := changer.Replace(univ)
	// changed := changer.Replace("Go$ Inha$") //"Go! Inha!로 변경 후 changed에 넣기"
	// fmt.Println(changed)                    //"Go! Inha!"

	// r := bufio.NewReader(os.Stdin) //버퍼를 가진 입력 리더 생성
	// // i, _ := r.ReadString('\n') //ignore error
	// i, err := r.ReadString('\n') //엔터까지 문자열을 읽음
	// // fmt.Println(err)
	// if err != nil { //입력 오류 발생 시 종료
	// 	log.Fatal(err) // report the error and exit the program
	// }
	// fmt.Println(i)

	// //shadowing
	// var float64 float64 = 2.71
	// var f float64 = 3.991
	// fmt.Println(float64)
	// fmt.Println(f)

	//var fmt float64 = 2.71
	//fmt.Println(fmt)	//패키지 이름과 같은 이름으로 선언 할 경우 오류

	// fmt.Print("Enter a score: ")
	// r := bufio.NewReader(os.Stdin)
	// i, err := r.ReadString('\n') //엔터가 입력될 때 까지 입력받음
	// if err != nil {              //입력 오류 발생 시 종료
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
	//현재 시간을 time.Time 타입으로 반환 후 1970년 1월 1일 이후 경과한 초를 int64 로 반환
	rand.Seed(seconds) //난수 생성기 시드 설정
	target := rand.Intn(100) + 1
	//0부터 99까지 랜덤 정수 반환 +1(1~100사이) / rand.Intn시드가 설정되지 않으면 항상 같은 수를 반환
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)         //표준 입력을 버퍼링된 Reader로 생성
	success := false                            //성공 여부 체크 변수(맞추면 true로 변경)
	for guesses := 0; guesses < 10; guesses++ { //10번 반복
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n') //엔터가 나올때 까지 입력
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //문자열 주위에 붙은 공란 및 탭 키 등 제거
		guess, err := strconv.Atoi(input) //문자열을 정수로 변환
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
