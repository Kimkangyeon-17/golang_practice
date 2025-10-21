package main

import "fmt"

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true  // 찾으면 즉시 반환
		}
	}
	return 0, false  // 못 찾으면 false 반환
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}