package main

import "strconv"

func fizzbuzz(n int) string {
	got := ""
	got += map[bool]string{true: "Fizz"}[n == 3]
	got += map[bool]string{true: strconv.Itoa(n)}[got == ""]
	return got
}
