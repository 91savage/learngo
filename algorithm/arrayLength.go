package main

import (
	"fmt"
)

func solution(strlist []string) []int {
	var result []int
	for _, v := range strlist {
		result = append(result, len(v))
	}

	return result
}

func main() {
	fmt.Println(solution([]string{"We", "are", "the", "world!"}))
}
