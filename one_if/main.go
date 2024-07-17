package main

import "strconv"

func fizzbuzz(n int) string {
	got := ""
	got += map[bool]string{true: "Fizz"}[n%3 == 0]
	if got == "" {
		got += strconv.Itoa(n)
	}
	return got
}
