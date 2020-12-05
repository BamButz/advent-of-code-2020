package inputs

import (
	"bufio"
	"log"
	"os"
)

func GetScanner(path string) *bufio.Scanner {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(bufio.NewReader(file))
	return scanner
}
