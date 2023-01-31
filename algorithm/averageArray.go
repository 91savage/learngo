package main

import (
	"fmt"
)

func solution(numbers []int) float64 {

	var tmp float64 =0
	i:=0
	var answer float64 = 0

	for i=0; i<len(numbers); i++ {
		tmp += float64(numbers[i])
	}
	answer = tmp / float64(len(numbers))
	return answer
}

func main(){
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}