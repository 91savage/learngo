package main

import (
	"fmt"
)

func solution(dot []int) int {

	switch {
	case dot[0] > 0 && dot[1] > 0:
		return 1
	case dot[0] < 0 && dot[1] > 0:
		return 2
	case dot[0] < 0 && dot[1] < 0:
		return 3
	default:
		return 4
	}
}

func main() {
	fmt.Println(solution([]int{2, 4}))
}

// ++ == 1
// -+ == 2
// -- == 3
// +- == 4

//dot[0] == x, dot[1] == y
