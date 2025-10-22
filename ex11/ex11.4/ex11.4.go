package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	fmt.Println("=== 복사 전 ===")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	b = a  // 배열 전체가 복사됨

	fmt.Println("\n=== 복사 후 ===")
	fmt.Println("b:", b)

	// a를 수정해도 b는 영향 없음
	a[0] = 999
	fmt.Println("\n=== a[0] 수정 후 ===")
	fmt.Println("a:", a)
	fmt.Println("b:", b)  // b는 변경되지 않음
}