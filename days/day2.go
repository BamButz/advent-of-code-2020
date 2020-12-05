package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code-2020/inputs"
)

func Two() {
	combinations := getPasswordAndPolicyCombinations()

	countOfValidPasswords := 0
	for _, combination := range combinations {
		if combination.isValid() {
			countOfValidPasswords++
		}
	}

	fmt.Println("How many passwords are valid according to their policies?")
	fmt.Printf("The answer is %d\n", countOfValidPasswords)

	countOfValidPasswords = 0
	for _, combination := range combinations {
		if combination.isValid2() {
			countOfValidPasswords++
		}
	}

	fmt.Println("How many passwords are valid according to the new interpretation of the policies?")
	fmt.Printf("The answer is %d\n", countOfValidPasswords)
}

func getPasswordAndPolicyCombinations() []*passwordPolicyCombination {
	scanner := inputs.GetScanner("inputs/day2.txt")
	regex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w*)$`)

	var passwords []*passwordPolicyCombination
	for scanner.Scan() {
		groups := regex.FindStringSubmatch(scanner.Text())

		password := groups[4]
		policy := passwordPolicy{}
		policy.min, _ = strconv.Atoi(groups[1])
		policy.max, _ = strconv.Atoi(groups[2])
		policy.char = groups[3]

		passwords = append(passwords, &passwordPolicyCombination{&policy, password})
	}

	return passwords
}

type passwordPolicy struct {
	min  int
	max  int
	char string
}

type passwordPolicyCombination struct {
	policy   *passwordPolicy
	password string
}

func (c *passwordPolicyCombination) isValid() bool {
	count := strings.Count(c.password, c.policy.char)

	if count < c.policy.min || count > c.policy.max {
		return false
	}

	return true
}

func (c *passwordPolicyCombination) isValid2() bool {
	if len(c.password) < c.policy.min {
		return false
	}

	firstPosition := false
	if string(c.password[c.policy.min-1]) == c.policy.char {
		firstPosition = true
	}

	if len(c.password) < c.policy.max {
		return firstPosition
	}

	if string(c.password[c.policy.max-1]) == c.policy.char {
		return !firstPosition
	}

	return firstPosition
}
