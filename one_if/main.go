package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Go")
}

func fizzbuzz(input int) string {
	got := ""
	got += map[bool]string{true: "Fizz"}[input%3 == 0]
	// got += map[bool]string{true: "Buzz"}[input%5 == 0]
	if got == ""{
		got = strconv.Itoa(input)
	}
	return got
}