package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("좋은 날씨가 아닙니다.")

	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있습니다.")

	case temp >= 15 && temp < 25:
		fmt.Println("좋은 날씨입니다.")

	default:
		fmt.Println("따뜻합니다.")
	}
}