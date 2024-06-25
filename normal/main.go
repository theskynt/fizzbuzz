package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Go")
}

func fizzbuzz(input int) string {
	// if input%3 == 0 && input%5 == 0 {
	// 	return "FizzBuzz"
	// }
	// if input%3 == 0 {
	// 	return "Fizz"
	// }
	// if input%5 == 0 {
	// 	return "Buzz"
	// }
	return strconv.Itoa(input)
}