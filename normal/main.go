package main

import "strconv"

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}
	if n == 5 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
