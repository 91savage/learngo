package main

import (
	"fmt"
)

func solution(numbers []int, k int) int {

	tmp := 1

	for i := 1; i < k; i++ {
		tmp += 2
		if tmp > len(numbers) {
			tmp = tmp % len(numbers)
		}

	}
	return tmp

}

func main() {
	fmt.Println(solution([]int{3, 1, 2}, 6))
}

//1,3,5,7->2,4,6->1,3,5
