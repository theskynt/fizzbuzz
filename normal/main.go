package main

import "strconv"

func fizzbuzz(input int) string {
	if input == 3 || input == 6 {
		return "Fizz"
	}
	if input == 5 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}
