package main

import (
	"fmt"
)

func solution(n int) int {

var discount float64 = 1
switch {
case n >= 500000:
	discount = 0.80
case n >= 300000:
	discount = 0.90
case n >= 100000:
	discount = 0.95
}
return int(float64(n)* discount)
}

func main(){
	fmt.Println(solution((150000)))
}


