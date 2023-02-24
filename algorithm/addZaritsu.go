package main

import (
	"fmt"
)

func solution(n int) int {

	// numStr := strconv.Itoa(n)
	// numArr := make([]int, len(numStr))
	// for i := 0; i < len(numStr); i++ {
	// 	num, _ := strconv.Atoi(string(numStr[i]))
	// 	numArr[i] = num
	// }
	// tmp := 0
	// for _, v := range numArr {
	// 	tmp = tmp + v
	// }
	// return tmp
	res := 0
	for n > 0 {
		// res = res + n%10
		// n = n / 10
		res = res + n%10
		n = n / 10
		fmt.Println(n)
	}
	return res

}

func main() {
	fmt.Println(solution(1234))
}
