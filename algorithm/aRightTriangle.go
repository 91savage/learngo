package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

// func main() {
// 	fmt.Println(solution(3))
// }
