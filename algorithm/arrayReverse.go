package main

import (
	"fmt"
)

func solution(num_list []int) []int{
	var reverse [] int
	len := len(num_list) -1
	
	for i:=len; i>=0; i-- {	
		reverse = append(reverse, num_list[i])
	}
	return reverse
}

func main(){
	fmt.Println(solution([]int{1,2,3,4,5}))
}