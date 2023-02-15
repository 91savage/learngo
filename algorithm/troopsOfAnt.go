package main

import (
	"fmt"
)

func solution(hp int) int {

	answer := 0

	if hp/5 != 0 {
		answer += hp / 5
		hp %= 5
	}
	if hp/3 != 0 {
		answer += hp / 3
		hp %= 3
	}
	if hp/1 != 0 {
		answer += hp / 1
	}
	return answer

}

func main() {
	fmt.Println(solution(22))
}
