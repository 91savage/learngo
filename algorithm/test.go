package main

import (
	"fmt"
)

func main() {
	const (
		Red int = iota
		Blue int = iota
		Yellow int = iota
	)
	fmt.Println(Red, Blue, Yellow)	
	// 0, 1, 2

	const (
		C1 uint = iota + 1
		C2
		C3
	)
	fmt.Println(C1,C2,C3)
}