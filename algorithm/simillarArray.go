package main

import (
	"fmt"
)

func solution(s1 []string, s2 []string) int {
	count := 0
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(solution([]string{"a", "b", "c"}, []string{"com", "b", "d", "p", "c"}))
}
