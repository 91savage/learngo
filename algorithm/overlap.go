package main

import (
	"fmt"
)

func solution(array []int, n int) int {
	count := 0
	for _, v := range array {
		if v == n {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(solution([]int{1, 1, 3, 4, 5}, 1))
}
