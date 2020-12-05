package days

import (
	"fmt"
	"log"
	"strconv"

	"advent-of-code-2020/inputs"
)

func One() {
	numbers := getNumbers()

	fmt.Println("Find the two entries that sum to 2020; what do you get if you multiply them together?")
	twoEntries := findTwoEntriesThatSumTo(numbers, 2020)
	fmt.Printf("The answer is %d\n", twoEntries[0]*twoEntries[1])

	fmt.Println("What is the product of the three entries that sum to 2020?")
	threeEntries := findThreeEntriesThatSumTo(numbers, 2020)
	fmt.Printf("The answer is %d\n", threeEntries[0]*threeEntries[1]*threeEntries[2])
}

func findThreeEntriesThatSumTo(numbers []int, sum int) []int {
	for _, summand := range numbers {
		otherSummands := findTwoEntriesThatSumTo(numbers, sum-summand)

		if otherSummands != nil {
			return append(otherSummands, summand)
		}
	}

	return nil
}

func findTwoEntriesThatSumTo(numbers []int, sum int) []int {
	summands := make(map[int]bool)
	for _, summand := range numbers {
		otherSummand := sum - summand

		if summands[otherSummand] {
			return []int{summand, otherSummand}
		}

		summands[summand] = true
	}

	return nil
}

func getNumbers() []int {
	scanner := inputs.GetScanner("inputs/day1.txt")

	var numbers []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}
