package main

import (
	"fmt"
)

func solution(angle int) int {

	switch {
	case angle < 90:
		return 1
	case angle == 90:
		return 2
	case angle < 180:
		return 3
	case angle == 180:
		return 4
	default:
		return -1
	}
}

func main() {
	fmt.Println(solution(91))

}
