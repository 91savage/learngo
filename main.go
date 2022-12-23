package main

import (
	"fmt"
	"strings"
)

// func repeatMe(words ...string) {
// 	fmt.Println(words)

// }

// naked return 
func lenAndUpper(name string) (length int, uppercase string) {
	// defer = return 이 끝나고 실행 됨.
	defer fmt.Println("I'm Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
		return
	// return length, uppercase 
}

// func lenAndUpper(name string) (int, string) {
//   // multi 값 return
// 	return len(name), strings.ToUpper(name)
// }


func multiply(a ,b int) int {  // 마지막 int는 return 할 값의 Type
	return a * b
}	

// func main() {
// 	repeatMe("nico", "peter", "Merry", "rudolph", "Saavy")
// }

func main() {
	totalLength, up := lenAndUpper("Savage")
	fmt.Println(totalLength, up)
	
}