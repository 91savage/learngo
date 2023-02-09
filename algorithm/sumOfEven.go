package main

import "fmt"

func solution(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(solution(10))
}
