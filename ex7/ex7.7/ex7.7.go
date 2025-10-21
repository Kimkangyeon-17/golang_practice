package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {  // 종료 조건
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)  // 재귀 호출
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("fibonacci(%d) = %d\n", i, fibonacci(i))
	}
}