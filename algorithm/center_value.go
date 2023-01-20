package main

import (
	"fmt"
    "sort"
)

func solution(array []int) int {
    sort.Ints(array)
    length := len(array) / 2 // array 의 길이
    return array[length]
}
func main(){
	fmt.Println(solution([]int{3,6,2,1,9,10}))
}
