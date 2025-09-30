package main

import (
	"fmt"
)

func main(){
	a := 3 // int64
	var b float64 = 3.5

	var c int = int(b)

	//d := a * b // a는 int, b는 float64이므로 타입이 달라서 에러 발생
	d := float64(a) * b

	var e int64 = 7
	//f := a * e // a는 int이고 e는 int64이므로 서로 int타입이더라도 타입이 달라서 에러 발생
	f := a * int(e)
	
	fmt.Println(a, b, c, d, e, f)
}