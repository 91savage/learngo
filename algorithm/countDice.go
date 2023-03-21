package main

import (
	"fmt"
)

func solution(box []int, n int) int {

	var tmp int = 1
	var answer int = 1
	for _, v := range box {
		tmp = v / n
		answer = answer * tmp
	}
	return answer
}

func main() {
	fmt.Println(solution([]int{10, 8, 6}, 3))
}
