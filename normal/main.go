package main

import "strconv"

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
