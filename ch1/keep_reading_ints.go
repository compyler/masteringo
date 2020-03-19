package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "END"{
			break // stop reading input
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			continue // skip this entry
		}
		sum += i
	}
	fmt.Println("All input sum:", sum)
}
