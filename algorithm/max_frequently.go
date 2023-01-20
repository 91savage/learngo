package main

import (
	"fmt"
)
// 각 숫자들이 등장한 횟수를 저장 할 맵 생성
func solution(array []int) int {
    // 각 숫자들이 등장한 횟수를 저장 할 맵 생성
    numMap := make(map[int]int)
    for _, num := range array {
        numMap[num]++
    }
    
    // 최빈값을 저장할 변수 선언
    mode := -1
    // 최빈값이 등장한 횟수를 저장할 변수 선언
    maxCount := 0
    // 최빈값이 등장한 횟수를 저장할 변수 선언
    modeCount := 0
    // 맵을 순회하면서 최빈값을 구함
    for num, count := range numMap {
        if count > maxCount {
            mode = num
            maxCount = count
            modeCount = 1
        } else if count == maxCount {
            modeCount++
        }
    }
    // 최빈값이 여러개면 -1 반환
    if modeCount > 1 {
        return -1
    }
    // 최빈값 반환
    return mode
}

func main(){
	fmt.Println(solution([]int{1, 2, 3, 3, 3, 4}))
}
