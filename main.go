package main

import (
	"fmt"
	"time"
)

// Goroutines : 다른 함수와 동시에 실행시키는 함수
// main 함수가 끝나면 Goroutines도 종료 됨.
// Main 함수는 결과를 '저장' 하는 곳.


func main() {
	c := make(chan string)
	people := [5]string{"Sehun", "Jins", "Merry","Mozzi","Ddung"}
	for _, person := range people {
		go isSexy(person, c) // 인수로 person(value값) 과 c(채널)을 받음
	}
	for i:=0; i<len(people); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is nice"
}
