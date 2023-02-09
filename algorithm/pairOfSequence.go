package main

import (
	"fmt"
)
func solution(n int) int {
    
	count :=0;

	// for i:=1; i<=n; i++ {
	// 	for j:=n; j>=1; j-- {
	// 		if i*j == n {
	// 			count++
	// 		}
	// 	}
	// }
	for i :=1; i<=n; i++ {
		if n % i == 0 {
			count++
		}
	}
	return count
}



func main(){
	fmt.Println(solution(20))
}
