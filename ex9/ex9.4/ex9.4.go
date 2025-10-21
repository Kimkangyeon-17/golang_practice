package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업이 있습니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습이 있습니다.")
	}
}