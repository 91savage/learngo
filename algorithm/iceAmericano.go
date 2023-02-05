package main

import (
	"fmt"
)

func solution(n int) []int {
	americano := 5500
	cup := n/americano
	change := n % americano

	var result = []int{cup, change}

	return result
}

func main() {
	fmt.Println(solution(5500))
}