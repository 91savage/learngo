package main

import (
	"fmt"
)

func solution(n int) []int {
	s_num :=1
	var arr []int
	var arr2 []int

	for s_num <= n {
		arr = append(arr, s_num)
		s_num++
	}
	answer := arr
// key , index
	for _, i := range answer {
		if i %2 != 0 {
		arr2 = append(arr2, i)
		}
	}
	return arr2

}



func main(){
	fmt.Println(solution(15))
}