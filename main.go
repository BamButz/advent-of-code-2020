package main

import (
	"fmt"
)

func main() {
	var dayRoutines []func()
	dayRoutines = append(
		dayRoutines,
	)

	for i, routine := range dayRoutines {
		fmt.Printf("### DAY %d ###\n", i+1)
		routine()
		fmt.Println()
	}
}
