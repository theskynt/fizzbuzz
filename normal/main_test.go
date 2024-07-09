package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("[1] input 1 return 1", func(t *testing.T) {
		input := 1
		want := "1"

		got := fizzbuzz(input)

		if got != want {
			t.Errorf("Want : %v but got %v", want, got)
		}
	})
}
