package main

import "fmt"

func main() {
	var a int16 = 3456 // a는 int 16타입 - 2byte 정수
	var b int8 = int8(a) // int16 -> int8

	fmt.Println(a, b)
}