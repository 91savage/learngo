package main

import (
	"fmt"
)

func solution(array []int, height int) int {
	result := 0
	for _, v := range array {
		if v > height {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(solution([]int{149, 180, 192, 170}, 167))
}
