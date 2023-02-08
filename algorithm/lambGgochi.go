package main

import (
	"fmt"
)


func solution(n,k int) int{
	
	ggochi:= 12000
	drink := 2000
	service := n/10

	if(n >= 10) {
		ggochi = ggochi * n
		drink = drink * k - (drink * service)
	}else{
		ggochi = ggochi * n
		drink = drink * k
	}


	return ggochi + drink

}


func main() {
	fmt.Println(solution(10,3))
}
