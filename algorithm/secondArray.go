package main

import (
	"fmt"
)

func solution(num_list []int, n int) [][]int {
	len := len(num_list) / n
	answer := make([][]int, len)

	for i := 0; i < len; i++ {
		answer[i] = make([]int, n)
		for j := 0; j < n; j++ {
			answer[i][j] = num_list[i*n+j]
		}
	}
	return answer
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))
}
