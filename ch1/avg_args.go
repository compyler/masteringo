package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	sum := 0.0
	count := 0
	for _, arg := range args {
		i, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			continue
		}
		sum += i
		count++
	}
	avg := sum / float64(count)

	fmt.Println("Average of all float arguments:", avg)
}
