package main

import (
	"fmt"
	"sort"
)

func solution(sides []int) int {
	sort.Slice(sides, func(i int, j int) bool {
		return sides[i] < sides[j]
	})
	if sides[0]+sides[1] > sides[2] {
		return 1
	} else {
		return 2
	}
}

func main() {
	fmt.Println(solution([]int{3, 6, 2}))
}
