package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn() 호출됨, cnt =", cnt)
	cnt++
	return cnt
}

func main() {
	// false && ... : && 왼쪽이 false이므로 오른쪽 평가 안 함
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 increase")
	}

	// true || ... : || 왼쪽이 true이므로 오른쪽 평가 안 함
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 increase")
	}

	fmt.Println("최종 cnt:", cnt)
}