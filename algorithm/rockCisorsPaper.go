package main

import (
	"fmt"
)

func solution(rsp string) string {

	answer := ""

	arr := make([]string, len(rsp))

	for i, s := range rsp {
		arr[i] = string(s)
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == "2" {
			answer += "0"
		} else if arr[i] == "0" {
			answer += "5"
		} else {
			answer += "2"
		}
	}
	return answer
}

func main() {
	fmt.Println(solution("2"))
}

//가위 2
//바위 0
//보 5
