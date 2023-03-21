package main

import (
	"fmt"
)

func solution(box []int, n int) int {

	tmp := 1
	for _, v := range box {
		tmp *= v / n

	}
	return tmp
}

func main() {
	fmt.Println(solution([]int{10, 8, 6}, 3))
}
