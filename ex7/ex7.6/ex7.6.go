package main

import "fmt"

func printNo(n int) {
	if n == 0 {  // 종료 조건 (base case)
		return
	}
	fmt.Println(n)
	printNo(n - 1)  // 재귀 호출
	fmt.Println("After", n)
}

func main() {
	printNo(3)
}