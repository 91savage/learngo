package main

import (
	"fmt"
)


func solution(hp int)int{
	
	big :=5
	middle :=3
	small :=1

	result :=0
	tmp :=0
	tmp2 :=0

	switch {
	case hp % big == 0:
		result = hp/big
	case hp % big != 0:
		tmp = hp % big
		if tmp >= 3 {
			tmp2 = tmp%middle
			result = hp/big+ tmp/middle + tmp2/small
		}else {
			tmp2 = tmp%middle
			result = hp/big + tmp2/small
		}

	}
	return result

}

func main(){
	fmt.Println(solution(22))
}

//[공격력]
//장군 5
//병정 3
//일 1

