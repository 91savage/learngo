package main

import (
	"fmt"
)

func solution(message string) int {
	return len(message) * 2
}

func main() {
	fmt.Println(solution("happy birthday!"))
}
