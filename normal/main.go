package main

import "strconv"

func fizzbuzz(input int) string {
	if input == 3 {
		return "Fizz"
	}
	if input%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(input)
}
