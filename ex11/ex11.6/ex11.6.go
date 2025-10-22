package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a [5]int        // int는 8바이트 (64비트 시스템)
	var b [10]float64   // float64는 8바이트
	var c [2][5]int     // [5]int가 2개

	fmt.Printf("[5]int 크기: %d바이트\n", unsafe.Sizeof(a))
	fmt.Printf("[10]float64 크기: %d바이트\n", unsafe.Sizeof(b))
	fmt.Printf("[2][5]int 크기: %d바이트\n", unsafe.Sizeof(c))
}