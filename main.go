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
	return strconv.Itoa(input)
}