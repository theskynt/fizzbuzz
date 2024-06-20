package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Go")
}

func fizzbuzz(input int) string {
	result := ""
    result += map[bool]string{true: "Fizz"}[input%3 == 0]
    result += map[bool]string{true: "Buzz"}[input%5 == 0]
	return map[bool]string{true: result, false: strconv.Itoa(input)}[result != ""]
}