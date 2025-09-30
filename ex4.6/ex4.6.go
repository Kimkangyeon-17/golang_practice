package main

import "fmt"

var g int = 10 // 패키지 전역 변수

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// m = s + 20 // s가 변수의 범위가 벗어나서 에러 발생

} 