package main

import (
	"fmt"
)

func solution(n int) int {
	cnt := 0

	for i := 4; i <= n; i++ {
		isComposite := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isComposite = true
				break
			}
		}
		if isComposite {
			cnt++
		}
	}
	return cnt

}

func main() {
	fmt.Println(solution(10))
}
