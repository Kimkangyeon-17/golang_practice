package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second)  // 1초 대기
		fmt.Println(i)
		i++
	}
}