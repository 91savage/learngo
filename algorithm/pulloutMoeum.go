package main

import (
	"fmt"
	"strings"
)

func solution(my_string string) string {
	split := strings.Split(my_string, "")
	var tmp []string

	fmt.Println(split)

	for _, v := range split {
		if v != "i" && v != "e" && v != "a" && v != "o" && v != "u" {
			tmp = append(tmp, v)
		}
	}
	answer := strings.Join(tmp, "")

	return answer
}

func main() {
	fmt.Println(solution("bus"))
}
