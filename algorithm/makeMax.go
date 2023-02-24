package main

import (
	"fmt"
	"sort"
)

func solution(numbers []int) int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	return numbers[0] * numbers[1]

}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5}))
}
