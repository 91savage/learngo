package main

import (
	"fmt"
	"strings"
)

func solution(my_string string) string {

	len := len(my_string) - 1
	var tmp []string

	for i := len; i >= 0; i-- {
		tmp = append(tmp, string(my_string[i]))

	}
	return strings.Join(tmp, "")

}

func main() {
	fmt.Println(solution("jaron"))

}
