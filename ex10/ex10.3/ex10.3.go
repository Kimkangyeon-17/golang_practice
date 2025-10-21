package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")
			stdin.ReadString('\n')
			continue  // 에러 발생 시 다음 반복으로
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break  // 짝수면 for 문 종료
		}
	}
	fmt.Println("for문이 종료됐습니다.")
}