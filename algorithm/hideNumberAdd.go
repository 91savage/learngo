package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(my_string string) int {
	count := 0
	str := strings.Split(my_string, "")

	for _, v := range str {
		i, _ := strconv.Atoi(v)
		count += i
	}
	return count
}

func main() {
	fmt.Println(solution("aAb1B2cC34oOp"))
}
