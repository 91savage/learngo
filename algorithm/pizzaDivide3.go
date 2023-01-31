package main

import (
	"fmt"
)

func solution(slice int, n int) int {

	tmp := n / slice 

		if n % slice == 0 {
			return tmp
		}else {
			return tmp +1
		}
		
}

func main(){
	fmt.Println(solution(7,10))
}