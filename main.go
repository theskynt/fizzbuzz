package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Go")
}

func fizzbuzz(input int) string {
	if input == 3 {
		return "Fizz"
	}

	if input == 5 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}