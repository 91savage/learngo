package main

import "fmt"

func solution(my_string string, letter string) string {

	var result string
	for i := 0; i < len(my_string); i++ {

		if string(my_string[i]) != letter {
			result += string(my_string[i])
		}
	}
	return result
}

func main() {
	fmt.Println(solution("abcdef", "f"))
}
