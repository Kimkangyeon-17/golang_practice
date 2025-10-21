package main

import "fmt"

func main() {
	for dan := 2; dan <= 9; dan++ {
		fmt.Printf("==== %d단 ====\n", dan)
		for num := 1; num <= 9; num++ {
			fmt.Printf("%d * %d = %d\n", dan, num, dan*num)
		}
		fmt.Println()
	}
}