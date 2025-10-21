package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {  // 초기문; 조건문; 후처리
		fmt.Print(i, ", ")      // i 값을 출력합니다.
	}
	// fmt.Println(i)  // 에러! i는 for 블록 밖에서 사용 불가
}