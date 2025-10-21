package main

import "fmt"

func main() {
	a := 1
	b := 1
	found := false

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break  // 안쪽 for만 빠져나옴
			}
		}
		if found {
			break  // 바깥쪽 for를 빠져나옴
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}