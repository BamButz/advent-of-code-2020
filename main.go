package main

import (
	"fmt"

	"advent-of-code-2020/days"
)

func main() {
	var dayRoutines []func()
	dayRoutines = append(
		dayRoutines,
		days.One,
	)

	for i, routine := range dayRoutines {
		fmt.Printf("### DAY %d ###\n", i+1)
		routine()
		fmt.Println()
	}
}
