package main

import (
	"fmt"
)

func solution(balls int, share int) int {
	answer := 1
	for i := 0; i < share; i++ {
		answer *= balls - i
		answer /= i + 1
	}
	return answer
}

func main() {
	fmt.Println(solution(3, 2))
}

// 전체 n개 중에서 a개를 순서없이 뽑는 경우의 수
//n * (n-1) * (n -2) 를 a개 만큼 진행한 값(분자) 나누기
//a * (a - 1) * (a - 2) 를 a가 1이 될때까지 진행한 값(분모)
