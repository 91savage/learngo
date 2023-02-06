package main

import "fmt"

func solution(n int) string {
	var result string

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			result += "*"
		}
		result += "\n"
	}
	return result
}

func main() {
	fmt.Println(solution(3))
}
