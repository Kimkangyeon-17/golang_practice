package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50}

	// 읽기
	fmt.Println(nums[0])  // 10
	fmt.Println(nums[2])  // 30

	// 쓰기
	nums[2] = 300

	// len() 함수로 배열 길이 확인
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}