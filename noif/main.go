package main

import "strconv"

func fizzbuzz(n int) string {
	got := ""
	got += map[bool]string{true: "Fizz"}[n%3 == 0]
	got += map[bool]string{true: "Buzz"}[n == 5]
	got += map[bool]string{true: strconv.Itoa(n)}[got == ""]
	return got
}
