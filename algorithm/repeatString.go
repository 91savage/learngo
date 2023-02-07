package main

import "fmt"

func solution(my_string string, n int) string {

	var result string

	for i := 0; i < len(my_string); i++ {
		for j := 0; j < n; j++ {
			result += string(my_string[i])
		}
	}
	return result

}

func main() {
	fmt.Println(solution("hello", 3))
}
