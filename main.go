package main

import (
	"time"
	"fmt"
)

// Goroutines : 다른 함수와 동시에 실행시키는 함수
// main 함수가 끝나면 Goroutines도 종료 됨.
// Main 함수는 결과를 '저장' 하는 곳.


func main() {
	go sexyCount("SEhun")
	go sexyCount("merry")
	time.Sleep(time.Second * 5)

}

func sexyCount(person string) {
	for i:=0; i<10; i++ {
		fmt.Println(person, "is sexy" , i)
		time.Sleep(time.Second)
	}
}