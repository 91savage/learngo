package main

import (
	"fmt"
)

func solution(numbers []int, direction string) []int {

	if direction == "right" {
		return append(numbers[len(numbers)-1:], numbers[:len(numbers)-1]...)
		// tmp := numbers[:len(numbers)-1]
		// tmp = append([]int{len(numbers)}, tmp...)
		// return tmp
	} else {
		return append(numbers[1:], numbers[0])
		// tmp := numbers[1:]
		// tmp = append(tmp, numbers[0])
		// return tmp
	}
}

func main() {
	fmt.Println(solution([]int{1, 2, 3}, "right"))
}
