package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		if i%2 != 0 {
			fmt.Println(i, "*", i, "=", i*i)
		}

	}
}
