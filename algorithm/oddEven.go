package main

import "fmt"

// 짝수가 먼저 나오게
func solution(num_list []int) []int {
	odd := 0
	even := 0
	var result []int
	for i := 0; i < len(num_list); i++ {
		if num_list[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	result = append(result, even)
	result = append(result, odd)

	return result
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5}))
}
