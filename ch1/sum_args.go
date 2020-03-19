package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	sum := 0
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		sum += i
	}
	fmt.Println("Sum of all integer arguments:", sum)
}
