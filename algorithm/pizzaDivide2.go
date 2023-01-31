package main

import (
	"fmt"
)
func solution(n int) int {
    piece :=6
	i :=0

	for i < n{
		if piece % n != 0 {
			piece +=6
		}
		i++	
	}
	return piece/6
}



func main(){
	fmt.Println(solution(12))
}
