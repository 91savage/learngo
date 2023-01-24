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
	people := [2]string{"Sehun", "Jins"}
	for _, person := range people {
		go isSexy(person, c) // 인수로 person(value값) 과 c(채널)을 받음
	}
	//위에 까진 즉시 실행
	// c의 값을 받을 떄까지 멈춰있다가 실행 됨. 그래서 10초 뒤에 실행
	fmt.Println("Waiting for messages") // 즉시 실행됨. c(채널)에서 값을 받지 않아도 되기 때문에
	resultOne := <-c
	resultTwo := <-c
	// resultThree := <-c
	fmt.Println("Received this message:", resultOne)
	fmt.Println("Received this message:", resultTwo)
	// fmt.Println("Received this message:", resultThree) 3개가 되면 에러가 뜸. GO는 배열에 선언된 2개가 나올것을 미리 예상하고 있음.

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + "is nice"
}
