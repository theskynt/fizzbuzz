package main

import "strconv"

func fizzbuzz(input int) string {
	if input == 3 {
		return "Fizz"
	}
	return strconv.Itoa(input)
}
